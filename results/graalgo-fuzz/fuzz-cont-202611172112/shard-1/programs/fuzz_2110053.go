package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint = 2
    v2 := make(map[int]uint32)
    v3 := []string{"", "a", ("hello" + "hello")}
    var v4 int = 9223372036854775807
    var v5 float64 = (3.14159 - (1.5 / -2.25))
    v6 := make(map[uint32]uint32)
    var v7 float32 = (1.5 / -2.25)
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if true {
                v4 += (fi9 ^ fi8)
                v10 := GenS1{F0: v1, F1: 65535, F2: 391}
                if (-781 == -9223372036854775808) {
                    var v11 int = fi9
                    delete(v2, fi9)
                    _ = v11
                }
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    switch v4 {
    case 42:
        if true {
            if true {
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 int64 = 42
                    _ = fi12
                    _ = v13
                }
                var v14 int32 = 1
                _ = v14
            }
            v15 := GenS1{F0: (v1 << ((v1) & 63)), F1: 65535, F2: -824}
            var v16 float64 = v5
            _ = v15
            _ = v16
        }
    case -2147483648:
        v17 := GenS2{F0: "a"}
        _ = v17
    default:
        if (true || (!((^(v1)) >= 65535))) {
            switch (("hello" + ("hello" + "")) + ("hello" + ((("hello" + (((("\n\t" + (("hello" + "hello") + ("a" + ("\n\t" + "\n\t")))) + "") + "\n\t") + ((("a" + "") + (("a" + "") + "\n\t")) + ("\n\t" + "foo bar")))) + "a") + "hello"))) {
            case "hello":
                var v18 bool = (!(664 < 32768))
                _ = v18
            case "":
                v7 -= 0.0
            case "foo bar":
                var v19 int16 = -32768
                _ = v19
            }
            v4 += (v4 % ((-1) | 1))
            var v20 int8 = -128
            _ = v20
        }
    }
    v21 := GenS2{F0: "\n\t"}
    v22 := GenS1{F0: 692, F1: 65535, F2: 0}
    v2[v4] = 862
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if true {
        v4 += (v4 * (v4 & v4))
        var v23 uint64 = 859
        _ = v23
    }
    v7 += (v7 + (v7 - v7))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
