using System;
using System.Collections.Generic;

namespace DataStructure
{
    public class MyLinkedList<T>
    {
        public MyNode<T> Last { get; set; }

        public MyNode<T> First { get; set; }

        public int Count { get; }

        public MyNode<T> AddFirst(T value)
        {
            var newNode = new MyNode<T>(value);
            if (First.Equals(default))
            {
                First = newNode;
                return First;
            }
            newNode.Next = First;
            First = newNode;
            return First;
        }

        public MyNode<T> AddLast(T value)
        {
            var newNode = new MyNode<T>(value);
            if (Last.Equals(default))
            {
                Last = newNode;
                return First;
            }
            Last.Next = newNode;
            Last = newNode;
            return Last;
        }

        public void RemoveFirst()
        {
            First = First.Next == default ? default : First.Next;
        }

        public void RemoveLast()
        {
            Last.Clear();
        }

        public bool Contains(T value)
        {
            var currentNode = First;
            while (currentNode.Next != default)
            {
                if (currentNode.Value.Equals(value))
                {
                    return true;
                }
            }
            return false;
        }

        public void InsertTo(int index, T value)
        {
            var currentNode = First;
            var preCurrentNode = new MyNode<T>();
            var newNode = new MyNode<T>(value);
            if (index == 0)
            {
                newNode.Next = First;
                First = newNode;
                return;
            }
            for (int i = 0; i < index; i++)
            {
                preCurrentNode = currentNode;
                currentNode = currentNode.Next;
                if (i + 1 == index)
                {
                    currentNode.Next = newNode;
                    Last = newNode;
                }
                if (currentNode.Next.Equals(default) && i + 1 < index)
                {
                    throw new IndexOutOfRangeException();
                }
            }
            newNode.Next = currentNode;
            preCurrentNode.Next = newNode;
        }

        public void ReplaceTo(int index, T value)
        {
            var currentNode = First;
            var preCurrentNode = new MyNode<T>();
            var newNode = new MyNode<T>(value);
            if (index == 0)
            {
                newNode.Next = First.Next;
                First = newNode;
            }
            for (int i = 0; i < index; i++)
            {
                currentNode = currentNode.Next;
                preCurrentNode = currentNode;
                if (i + 1 == index)
                {
                    preCurrentNode.Next = newNode;
                    Last = newNode;
                }
                if (currentNode.Next.Equals(default) && i + 1 < index)
                {
                    throw new IndexOutOfRangeException();
                }
            }
            preCurrentNode.Next = newNode;
            newNode.Next = currentNode;
        }

        public void Reverse()
        {
            var tmp = new List<MyNode<T>>();
            var currentNode = First;
            while(currentNode.Next != default)
            {
                tmp.Add(currentNode);
                currentNode = currentNode.Next;
            }
            currentNode = new MyNode<T>();
            for (int i = tmp.Count-1; i >= 0; i--)
            {
                if (i == tmp.Count)
                {
                    First = tmp[i];
                    currentNode = First;
                    continue;
                }
                currentNode.Next = tmp[i];
            }
            Last = currentNode;
        }
    }

    public class MyNode<T>
    {
        public MyNode()
        { 
        
        }

        public MyNode(T value)
        {
            Value = value;
        }

        //public MyNode<T> Previous { get; set; }
    
        public MyNode<T> Next { get; set; }

        public T Value { get; set; }

        public void Clear()
        {
            Value = default(T);
            Next = default(MyNode<T>);
        }
    }
}
