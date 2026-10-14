package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 string; F2 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 243
    var v3 uint = (v2 | 2)
    v3 = 1
    if true {
        v4 := GenS1{F0: false, F1: ("foo bar" + "a"), F2: 100}
        if (59 < 127) {
            for fi5 := 0; fi5 < 12; fi5++ {
                switch ("foo bar" + ("foo bar" + (("foo bar" + "foo bar") + ("a" + "foo bar")))) {
                case "a":
                    v6 := GenS2{F0: 0}
                    _ = v6
                case "hello":
                    switch (fi5 % ((fi5) | 1)) {
                    case 9223372036854775807:
                        if ((-(fi5)) != 42) {
                            v7 := []int{(fi5 << ((730) & 63))}
                            switch "hello" {
                            case "\n\t":
                                v8 := GenS2{F0: 0}
                                _ = v8
                            case "a":
                                v3 = v3
                            }
                            _ = v7
                        }
                    }
                case "foo bar":
                    v9 := []int{fi5}
                    _ = v9
                }
                var v10 float64 = 1.5
                _ = fi5
                _ = v10
            }
        } else {
            var v11 uint16 = 2
            v12 := []int{42}
            v13 := GenS3{F0: (v3 + (v3 - 591)), F1: ("a" + (("\n\t" + ("a" + (("hello" + "hello") + ""))) + (("hello" + "\n\t") + "foo bar"))), F2: -1}
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 int8 = 42
        _ = v14
        _ = v4
    }
    fmt.Println(v2, v2)
    var v15 int = 9223372036854775807
    for fi16 := 0; fi16 < 5; fi16++ {
        v3 += (v3 << (((v3 & v2)) & 63))
        fmt.Println(v15)
        _ = fi16
    }
    v17 := GenS3{F0: v3, F1: "hello", F2: 97}
    var v18 uint = v2
    fmt.Println(v3, v15, v2)
    var v19 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
