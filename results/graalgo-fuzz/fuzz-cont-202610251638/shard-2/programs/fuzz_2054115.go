package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 int16; F2 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 int = -2147483648
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := []int{fi5, 735}
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    if false {
        if (0.0 > 1.5) {
            delete(v1, (v3 >> ((329) & 63)))
        }
    }
    v3 = (^(v3))
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: "a", F1: 18446744073709551615, F2: 2}
    var v9 uint64 = 18446744073709551615
    v10 := GenS1{F0: (((("foo bar" + "foo bar") + "\n\t") + ("" + ((("\n\t" + "a") + "\n\t") + ("foo bar" + "foo bar")))) + (("" + ("hello" + (("" + "hello") + ("hello" + "")))) + "\n\t")), F1: 4294967295, F2: 2}
    fmt.Println(v7)
    switch int(v7) {
    case -2147483648:
        v11 := GenS2{F0: 42, F1: 42, F2: 255}
        _ = v11
    case 639:
        _ = len(v2)
    default:
        for fi12 := 0; fi12 < 5; fi12++ {
            if (1 != 696) {
                v13 := []int{fi12, v3}
                _ = v13
            } else {
                var v14 uint16 = 65535
                v3 += v3
                _ = v14
            }
            v9 = (^(v9))
            switch "hello" {
            case "\n\t":
                switch (-(fi12)) {
                case 32768:
                    v9 -= (^(v9))
                case 816:
                    var v15 float64 = -2.25
                    _ = v15
                }
            case "a":
                v16 := GenS3{F0: 255, F1: -1, F2: 759}
                _ = v16
            }
            _ = fi12
        }
    }
    v17 := GenS1{F0: "hello", F1: 1, F2: v9}
    v18 := GenS1{F0: "", F1: uint(v9), F2: v9}
    v19 := []int{(v3 << ((18446744073709551615) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
