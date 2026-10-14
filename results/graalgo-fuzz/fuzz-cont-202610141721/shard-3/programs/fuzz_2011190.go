package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 0
    if (0 != 2147483647) {
        fmt.Println(v3)
    } else {
        var v4 string = ("\n\t" + "foo bar")
        v4 = v4
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := []int{fi5, 127}
            var v7 int = (fi5 % (((fi5 & (fi5 << ((42) & 63)))) | 1))
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = v4
    }
    v8 := GenS1{F0: 127}
    v3 = v3
    if true {
        if (uint32(v3) < 989) {
            for fi9 := 0; fi9 < 10; fi9++ {
                v3 -= 65535
                var v10 uint32 = 1
                _ = fi9
                _ = v10
            }
            v3 -= (v3 - (v3 * 2))
        } else {
            v3 = 731
            v11 := []int{-2147483648, -2147483648, -152}
            _ = v11
        }
        if true {
            v12 := v2[804]
            var v13 int32 = 852
            _ = v12
            _ = v13
        } else {
            var v14 int32 = 0
            for fi15 := 0; fi15 < 17; fi15++ {
                v3 -= (v3 | v3)
                var v16 uint8 = 1
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        fmt.Println(v3)
    }
    if (32767 < 439) {
        v17 := []int{2147483647, 0, 128}
        _ = v17
    }
    var v18 uint32 = 42
    fmt.Println(v3, v3)
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            v2[v18] = ("" + "foo bar")
            if (!true) {
                var v21 int16 = 42
                switch "foo bar" {
                case "hello":
                    v22 := GenS1{F0: 255}
                    _ = v22
                case "a":
                    var v23 bool = (!((!true) || false))
                    v24 := GenS1{F0: 42}
                    fmt.Println(v18, v23)
                    _ = v23
                    _ = v24
                }
                if ((0.0 == (3.14159 / (((0.0 + -2.25) / float64((3.14159 / (((3.14159 * float32(3.14159)) - (0.0 - 1e10)) + float32(1.5))))) - (-2.25 / -2.25)))) || (v21 > (v21 ^ (v21 - (v21 * (v21 * v21)))))) {
                    if true {
                        v1[fi19] = (true || (!((!true) || false)))
                        v25 := GenS2{F0: ((42 != 151) && true), F1: -128}
                        v26 := GenS2{F0: true, F1: -1}
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS3{F0: 32767, F1: 0, F2: int8(v21)}
                    delete(v2, (^(v18)))
                    _ = v27
                }
                _ = v21
            }
            v18 -= (^(v18))
            _ = fi20
        }
        if (204 < 348) {
            var v28 uint64 = 18446744073709551615
            _ = v28
        } else {
            fmt.Println(v18)
        }
        v29 := GenS1{F0: 0}
        _ = fi19
        _ = v29
    }
    var v30 int64 = -9223372036854775808
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := []int{(fi31 - int(v3)), (-(fi31)), (fi31 - 127), fi31}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        v34 := v1[fi33]
        var v35 float32 = 1e10
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: 108, F1: 2, F2: -128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v8
}
