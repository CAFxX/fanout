package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int32 }
type GenS3 struct { F0 int; F1 string }

func main() {
    v1 := []string{"a", ("a" + "foo bar"), ("" + ("a" + "\n\t")), ""}
    var v2 uint32 = 2147483648
    var v3 uint32 = 65535
    var v4 uint = 158
    v5 := make(map[string]int)
    var v6 int32 = -1
    if (631 == v2) {
        v7 := GenS1{F0: 819, F1: 42, F2: 127}
        v4 = v4
        _ = v7
    } else {
        v8 := v5["foo bar"]
        v4 += v4
        var v9 string = "hello"
        _ = v8
        _ = v9
    }
    v2 += 2147483648
    v10 := GenS1{F0: 18446744073709551615, F1: 717, F2: 128}
    var v11 bool = (!false)
    var v12 int = 1
    switch "foo bar" {
    case "hello":
        var v13 uint64 = 0
        _ = v13
    case "a":
        if ((-624 != 744) || (42 < 128)) {
            if v11 {
                v14 := v5["foo bar"]
                v12 = 1
                var v15 bool = false
                _ = v14
                _ = v15
            }
            var v16 int = -402
            _ = v16
        }
    default:
        v17 := GenS2{F0: 0, F1: -2147483648, F2: v6}
        _ = v17
    }
    v12 = (-(v12))
    var v18 float64 = 1.5
    fmt.Println(v18, v3, v11)
    var v19 string = (("foo bar" + ("\n\t" + "hello")) + ("" + ((("" + "\n\t") + ("hello" + ("" + (("hello" + ("a" + "a")) + ("\n\t" + "foo bar"))))) + ("\n\t" + "\n\t"))))
    if v11 {
        v18 = v18
        switch (v19 + v19) {
        case "hello":
            v11 = v11
        default:
            for fi20 := 0; fi20 < 18; fi20++ {
                v1 = append(v1, v19)
                for fi21 := 0; fi21 < 15; fi21++ {
                    v19 = ("" + v19)
                    v1 = append(v1, v19)
                    _ = fi21
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 int32 = v6
                    v6 = v23
                    fmt.Println(v23, v3, fi20)
                    switch "foo bar" {
                    case "":
                        v5[(v19 + "\n\t")] = v12
                    case "\n\t":
                        v3 = (v3 >> (((v4 + (^(v4)))) & 63))
                    }
                    v24 := GenS3{F0: 32767, F1: (((v19 + "\n\t") + v19) + ("\n\t" + "\n\t"))}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi20
            }
        }
    } else {
        var v25 string = (v19 + (v19 + v19))
        _ = v25
    }
    var v26 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
