
reStructuredText Markup
=========================================================================

See: `reStructuredText Markup Specification <https://docutils.sourceforge.io/docs/ref/rst/restructuredtext.html>`_

.. contents:: Table of Contents
  :depth: 3
  :local:
  :backlinks: none

Paragraphs
----------

Paragraphs consist of blocks of left-aligned text with no markup indicating any other body element. Blank lines separate paragraphs from each other and from other body elements. Paragraphs may contain inline markup.

Bullet Lists
------------

- This is the first bullet list item.  The blank line above the
  first list item is required; blank lines between list items
  (such as below this paragraph) are optional.

- This is the first paragraph in the second item in the list.

  This is the second paragraph in the second item in the list.
  The blank line above this paragraph is required.  The left edge
  of this paragraph lines up with the paragraph above, both
  indented relative to the bullet.

  - This is a sublist.  The bullet lines up with the left edge of
    the text blocks above.  A sublist is a new list so requires a
    blank line above and below.

- This is the third item of the main list.

This paragraph is not part of the list.

Enumerated Lists
----------------

1. Item 1 initial text.

   a) Item 1a.
   b) Item 1b.

2. a) Item 2a.
   b) Item 2b.

Definition Lists
----------------

term 1
    Definition 1.

term 2
    Definition 2, paragraph 1.

    Definition 2, paragraph 2.

term 3 : classifier
    Definition 3.

term 4 : classifier one : classifier two
    Definition 4.

Field Lists
-----------

:Date: 2001-08-16
:Version: 1
:Authors: - Me
          - Myself
          - I
:Indentation: Since the field marker may be quite long, the second
   and subsequent lines of the field body do not have to line up
   with the first line, but they must be indented relative to the
   field name marker, and they must line up with each other.
:Parameter i: integer

Bibliographic Fields
^^^^^^^^^^^^^^^^^^^^

Abstract 
~~~~~~~~~~~~~~~~

.. topic:: Title

   Body.

Address
~~~~~~~~~~~~~~~~

:Address: 123 Example Ave.
          Example, EX

Author
~~~~~~~~~~~~~~~~

:Author: J. Random Hacker

Authors
~~~~~~~~~~~~~~~~

:Authors: J. Random Hacker; Jane Doe

Contact
~~~~~~~~~~~~~~~~

:Contact: jrh@example.com

Copyright
~~~~~~~~~~~~~~~~

:Copyright: This document has been placed in the public domain.

Date
~~~~~~~~~~~~~~~~

:Date: 2002-08-20

Organization
~~~~~~~~~~~~~~~~

:Organization: Humankind

Version - Revision
~~~~~~~~~~~~~~~~~~

:Version: 1
:Revision: b

Status
~~~~~~~~~~~~~~~~

:Status: Work In Progress

RCS Keywords
^^^^^^^^^^^^

:Status: $keyword: expansion text $

:Status: Work In Progress

Option Lists
------------

-a         Output all.
-b         Output both (this description is
           quite long).
-c arg     Output just arg.
--long     Output all day long.

-p         This option has two paragraphs in the description.
           This is the first.

           This is the second.  Blank lines may be omitted between
           options (as above) or left in (as here and below).

--very-long-option  A VMS-style option.  Note the adjustment for
                    the required two spaces.

--an-even-longer-option
           The description can also start on the next line.

-2, --two  This option has two variants.

-f FILE, --file=FILE  These two options are synonyms; both have
                      arguments.

/V         A VMS/DOS-style option.

Literal Blocks
--------------

This is a typical paragraph.  An indented literal block follows.

::

    for a in [5,4,3,2,1]:   # this is program code, shown as-is
        print a
    print "it's..."
    # a literal block continues until the indentation ends

This text has returned to the indentation of the first paragraph,
is outside of the literal block, and is therefore treated as an
ordinary paragraph.

Indented Literal Blocks
-----------------------

This is a typical paragraph.  An indented literal block follows::

    for a in [5,4,3,2,1]:   # this is program code, shown as-is
        print a
    print "it's..."
    # a literal block continues until the indentation ends

This text has returned to the indentation of the first paragraph,
is outside of the literal block, and is therefore treated as an
ordinary paragraph.

Line Blocks
-----------

| Lend us a couple of bob till Thursday.
| I'm absolutely skint.
| But I'm expecting a postal order and I can pay you back
  as soon as it comes.
| Love, Ewan.

--------------

Take it away, Eric the Orchestra Leader!

    | A one, two, a one two three four
    |
    | Half a bee, philosophically,
    |     must, *ipso facto*, half not be.
    | But half the bee has got to be,
    |     *vis a vis* its entity.  D'you see?
    |
    | But can a bee be said to be
    |     or not to be an entire bee,
    |         when half the bee is not a bee,
    |             due to some ancient injury?
    |
    | Singing...

Block Quotes
------------

This is an ordinary paragraph, introducing a block quote.

    "It is my business to know things.  That is my trade."

    -- Sherlock Holmes

Doctest Blocks
--------------

This is an ordinary paragraph.

>>> print 'this is a Doctest block'
this is a Doctest block

The following is a literal block::

    >>> This is not recognized as a doctest block by
    reStructuredText.  It *will* be recognized by the doctest
    module, though!

Grid Tables
-----------

+------------------------+------------+----------+----------+
| Header row, column 1   | Header 2   | Header 3 | Header 4 |
| (header rows optional) |            |          |          |
+========================+============+==========+==========+
| body row 1, column 1   | column 2   | column 3 | column 4 |
+------------------------+------------+----------+----------+
| body row 2             | Cells may span columns.          |
+------------------------+------------+---------------------+
| body row 3             | Cells may  | - Table cells       |
+------------------------+ span rows. | - contain           |
| body row 4             |            | - body elements.    |
+------------------------+------------+---------------------+

+--------------+----------+-----------+-----------+
| row 1, col 1 | column 2 | column 3  | column 4  |
+--------------+----------+-----------+-----------+
| row 2        | Use the command ``ls | more``.   |
|              |                                  |
+--------------+----------+-----------+-----------+
| row 3        |          |           |           |
+--------------+----------+-----------+-----------+

Simple Tables
-------------

=====  =====  =======
  A      B    A and B
=====  =====  =======
False  False  False
True   False  False
False  True   False
True   True   True
=====  =====  =======

=====  =====  ======
   Inputs     Output
------------  ------
  A      B    A or B
=====  =====  ======
False  False  False
True   False  True
False  True   True
True   True   True
=====  =====  ======

=====  =====
col 1  col 2
=====  =====
1      Second column of row 1.
2      Second column of row 2.
       Second line of paragraph.
3      - Second column of row 3.

       - Second item in bullet
         list (row 3, column 2).
\      Row 4; column 1 will be empty.
=====  =====

Footnotes
---------

.. [1] Body elements go here.

If [#note]_ is the first footnote reference, it will show up as
"[1]".  We can refer to it again as [#note]_ and again see
"[1]".  We can also refer to it as note_ (an ordinary internal
hyperlink reference).

.. [#note] This is the footnote labeled "note".

Here is a symbolic footnote reference: [*]_.

.. [*] This is the footnote.

[2]_ will be "2" (manually numbered),
[#]_ will be "3" (anonymous auto-numbered), and
[#label]_ will be "1" (labeled auto-numbered).

.. [2] This footnote is labeled manually, so its number is fixed.

.. [#label] This autonumber-labeled footnote will be labeled "1".
   It is the first auto-numbered footnote and no other footnote
   with label "1" exists.  The order of the footnotes is used to
   determine numbering, not the order of the footnote references.

.. [#] This footnote will be labeled "3".  It is the second
   auto-numbered footnote, but footnote label "2" is already used.

Citations
---------

Here is a citation reference: [CIT2002]_.

.. [CIT2002] This is the citation.  It's just like a footnote,
   except the label is textual.

Hyperlink Targets
-----------------

Explicit Internal Targets
^^^^^^^^^^^^^^^^^^^^^^^^^

Clicking on this internal hyperlink will take us to the target_ below.

Clicking on this internal hyperlink will take us to the target1_ below.

Clicking on this internal hyperlink will take us to the target2_ below.

.. _target:

The hyperlink target above points to this paragraph.

.. _target1:
.. _target2:

The targets "target1" and "target2" are synonyms; they both point to this paragraph.

Implicit Internal Targets (Headings)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Jump to the `Bullet Lists`_ heading.

Other Document Topic Targets
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Go to the other part of this document, the `reStructuredText Directives <reSTDirectives.html>`_ chapter.

External Targets
^^^^^^^^^^^^^^^^^^^^^^^^^

See the Python_ home page for info.

`Write to me`_ with your questions.

.. _Python: http://www.python.org
.. _Write to me: jdoe@example.com

See the `Python home page <http://www.python.org>`_ for info.

This `link <Python home page_>`_ is an alias to the link above.

This is exactly equivalent to:

See the `Python home page`_ for info.

Comments
--------

.. This is a comment
..
   _so: is this!
..
   [and] this!
..
   this:: too!
..
   |even| this:: !

Inline Markup
-------------

*emphasis*

**strong emphasis**

:literal:`interpreted text` is the same as ``interpreted text``.



This is :sup:`superscript` interpreted text.

``inline literals``

