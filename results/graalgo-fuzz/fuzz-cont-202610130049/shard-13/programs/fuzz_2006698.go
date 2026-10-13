package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 2147483647
    v3 := []string{}
    _ = len(v1)
    v2 += (^(v2))
    fmt.Println(v2)
    fmt.Println(v2)
    switch ("a" + ((("" + ("foo bar" + ("" + "\n\t"))) + "foo bar") + "foo bar")) {
    case "hello":
        switch -38 {
        case 0:
            var v4 float32 = 0.0
            _ = v4
        }
    case "":
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 int64 = 42
            var v7 float64 = 3.14159
            if ((-94 < -56) && (-9223372036854775808 == fi5)) {
                v3 = append(v3, "hello")
                v8 := GenS1{F0: (42 < 255)}
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS2{F0: 1}
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v12 := GenS3{F0: 2}
                    fmt.Println(v7, fi11)
                    _ = len(v3)
                    _ = fi11
                    _ = v12
                }
                var v13 int8 = 0
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v13
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        if (0.1 == 1e10) {
            v2 = 0
        } else {
            v2 = v2
        }
    }
    v2 += (v2 + v2)
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS3{F0: 2}
        v2 -= (^(v2))
        v2 = (v2 | (v2 * v2))
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 42}
        v2 -= v2
        _ = fi16
        _ = v17
    }
    v18 := v1["hello"]
    if (!true) {
        v18 = (v18 && (4294967295 >= 128))
        var v19 int16 = -1
        v20 := GenS2{F0: 222}
        _ = v19
        _ = v20
    }
    var v21 int8 = 127
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int64 = -9223372036854775808
        _ = len(v3)
        v2 = (v2 | v2)
        _ = fi22
        _ = v23
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
