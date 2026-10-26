package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 uint, p2 bool) uint64 {
    var v6 float64 = -2.25
    p0 -= (p0 >> ((441) & 63))
    switch ("foo bar" + "hello") {
    case "foo bar":
        v7 := []int{-56, -2147483648}
        _ = v7
    case "a":
        var v8 int64 = -976
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 6; fi9++ {
            if p2 {
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := []int{32768, fi10, (fi10 - (fi10 * int(p1))), fi10}
                    var v12 bool = (p2 && p2)
                    if (!(p0 == (p1 % (((p0 - 0)) | 1)))) {
                        v6 += 3.14159
                    } else {
                        p1 += 799
                        p1 -= 664
                        v13 := GenS1{F0: 299, F1: 42}
                        var v14 int32 = 42
                        v12 = p2
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            var v15 uint16 = 65535
            if true {
                v16 := GenS2{F0: 2147483647}
                fmt.Println(v15)
                v6 = v6
                _ = v16
            } else {
                var v17 uint = 65535
                p1 = p0
                _ = v17
            }
            _ = fi9
            _ = v15
        }
    default:
        v6 += 1.5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 797
}

func main() {
    var v1 float32 = 0.1
    v2 := []string{("foo bar" + (("\n\t" + "hello") + ((("foo bar" + "hello") + ("hello" + "a")) + "a")))}
    v3 := []uint32{0}
    v4 := []int{-509, 502, 0}
    var v5 bool = false
    v5 = (-2147483648 < 1)
    var v18 uint64 = genFn0(1, 151, (v5 && v5))
    v1 += v1
    v5 = (-2147483648 != 127)
    v19 := GenS1{F0: -2147483648, F1: (^(v18))}
    switch "a" {
    case "foo bar":
        v4 = append(v4, -1)
    }
    if true {
        var v20 uint16 = 0
        v18 -= 0
        v1 = v1
        _ = v20
    }
    if (!v5) {
        v21 := GenS3{F0: v5}
        var v22 uint8 = 0
        if v5 {
            switch 1 {
            case -1:
                fmt.Println(v1, v5, v22)
            case 827:
                fmt.Println(v18, v18, v18)
            }
        }
        _ = v21
        _ = v22
    } else {
        v23 := GenS2{F0: -112}
        _ = v23
    }
    var v24 bool = v5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
