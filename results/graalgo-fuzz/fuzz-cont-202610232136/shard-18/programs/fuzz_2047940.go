package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 string; F2 bool }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 1
    var v2 int64 = 1
    var v3 int64 = (-(v2))
    for fi4 := 0; fi4 < 3; fi4++ {
        v1 = (v1 ^ (^(v1)))
        v1 = 18446744073709551615
        for fi5 := 0; fi5 < 11; fi5++ {
            if true {
                var v6 int64 = (v3 << ((4294967295) & 63))
                _ = v6
            }
            var v7 string = ("" + ("hello" + "\n\t"))
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    v8 := GenS1{F0: v1, F1: 224}
    v1 = (^(v1))
    if (v2 != (v3 + v3)) {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS2{F0: fi9, F1: ("\n\t" + (("a" + "\n\t") + "")), F2: false}
            _ = fi9
            _ = v10
        }
        if true {
            switch "\n\t" {
            case "foo bar":
                if ((-(v2)) > (v3 & (^(v2)))) {
                    var v11 uint32 = 894
                    v12 := GenS3{F0: 0}
                    _ = v11
                    _ = v12
                }
            case "hello":
                var v13 int64 = (v3 & (v3 / ((v2) | 1)))
                _ = v13
            default:
                v14 := []int{866, 9223372036854775807}
                _ = v14
            }
        }
        if (!(0.0 > (float64(0.1) - (1.5 - float64(-2.25))))) {
            var v15 int = 1
            var v16 uint64 = (v1 % ((v1) | 1))
            switch (v15 | v15) {
            case 127:
                v17 := GenS2{F0: v15, F1: "\n\t", F2: (v15 > 1)}
                _ = v17
            case 0:
                fmt.Println(v15, v16)
            default:
                var v18 int16 = 0
                _ = v18
            }
            _ = v15
            _ = v16
        } else {
            v1 = (^(v1))
        }
    }
    var v19 uint = 18446744073709551615
    v20 := GenS2{F0: 9223372036854775807, F1: ("a" + ""), F2: false}
    v2 += (v2 << ((v19) & 63))
    var v21 string = ("a" + ("a" + "a"))
    v2 -= v3
    v21 = ((v21 + (v21 + (("hello" + v21) + (v21 + "foo bar")))) + ("" + "a"))
    fmt.Println(v21, v21, v21)
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{(fi22 | fi22), fi22, 32768, (fi22 / (((fi22 & (fi22 | 0))) | 1))}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
}
