package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 bool) int64 {
    var v7 int = 32767
    switch "a" {
    case "hello":
        if p0 {
            var v8 uint64 = 4294967295
            if (!(-131 <= -737)) {
                var v9 bool = true
                v10 := []int{(-(v7)), v7, (v7 / ((v7) | 1))}
                _ = v10
                _ = v9
            }
            p0 = (false && (!(v8 < (^(v8)))))
            _ = v8
        } else {
            v11 := []int{(v7 - v7)}
            v12 := GenS1{F0: -1}
            v11 = append(v11, 491)
            _ = v11
            _ = v12
        }
    default:
        v13 := []int{v7, (v7 + v7), 32768, v7}
        _ = v13
    }
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 bool = (138 != 108)
    var v2 int = 513
    v3 := make(map[int]bool)
    v4 := []uint32{690, 0}
    var v5 int32 = 2147483647
    var v6 uint8 = 1
    var v14 float32 = 1e10
    if (-48 != 0) {
        var v15 bool = true
        _ = len(v3)
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        _ = v15
    }
    v5 += v5
    if v1 {
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                v1 = (("" + "hello") > "foo bar")
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi17
            }
            _ = fi16
        }
        v18 := GenS2{F0: ""}
        v6 = (v6 | v6)
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS1{F0: (v5 * (v5 << ((42) & 63)))}
        v1 = (!(22 != 65535))
        _ = fi19
        _ = v20
    }
    v4 = append(v4, 2147483647)
    for fi21 := 0; fi21 < 17; fi21++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi21
    }
    v22 := GenS2{F0: ("a" + "a")}
    if true {
        v14 = ((float32(((0.1 * 1.5) / (1e10 / (1.5 / -2.25)))) * v14) / 3.14159)
        switch "hello" {
        case "\n\t":
            v23 := GenS2{F0: ""}
            _ = v23
        }
    } else {
        fmt.Println(v14, v2)
        v24 := v3[v2]
        _ = v24
    }
    v6 += v6
    var v25 bool = false
    var v26 uint16 = 156
    v1 = ((!v1) || v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
