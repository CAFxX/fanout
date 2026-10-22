package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 int32) bool {
    v7 := []int{}
    var v8 uint16 = 255
    p1 -= p1
    if true {
        for fi9 := 0; fi9 < 19; fi9++ {
            v8 = 32767
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 string = ""
                v12 := GenS1{F0: 127}
                p0 += (-(p0))
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 int32 = p1
            _ = fi9
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1() int64 {
    var v14 float32 = (1e10 - -2.25)
    if false {
        v15 := GenS1{F0: -9223372036854775808}
        v14 -= float32(float64(float32(((0.0 + 0.0) + (3.14159 * 0.1)))))
        v14 -= 0.1
        _ = v15
    }
    v16 := []int{453, 9223372036854775807}
    if true {
        if len(v16) > 0 {
            v16[0] = 872
        }
        var v17 bool = false
        for fi18 := 0; fi18 < 16; fi18++ {
            v14 = 1.5
            v16 = append(v16, 719)
            _ = fi18
        }
        _ = v17
    } else {
        var v19 bool = (1000 != 1)
        _ = v19
    }
    _ = v14
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]bool)
    var v3 int8 = 90
    var v4 float64 = 1.5
    var v5 int64 = -1
    var v6 uint32 = 65535
    var v20 int8 = (v3 + v3)
    fmt.Println(v5)
    v21 := GenS1{F0: 127}
    v22 := []int{0}
    var v23 float64 = 1e10
    _ = len(v22)
    var v24 uint8 = 53
    v24 = (v24 | (v24 << ((88) & 63)))
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS1{F0: -374}
        _ = fi25
        _ = v26
    }
    v24 -= (v24 | v24)
    if true {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        fmt.Println(v3, v4)
    } else {
        switch 32767 {
        case -2147483648:
            _ = len(v2)
        case 127:
            fmt.Println(v6, v5)
        case -1:
            var v27 int32 = 42
            _ = v27
        }
        var v28 int32 = 820
        v5 = v1
        _ = v28
    }
    if false {
        if (false || false) {
            var v29 int16 = 1000
            _ = v29
        } else {
            fmt.Println(v6)
            var v30 uint8 = uint8(v6)
            if false {
                if len(v22) > 0 {
                    v22[0] = 900
                }
                v31 := GenS1{F0: 42}
                switch "foo bar" {
                case "a":
                    v3 += v3
                default:
                    var v32 string = "\n\t"
                    _ = v32
                }
                _ = v31
            }
            _ = v30
        }
        v33 := GenS2{F0: -408}
        v24 = (v24 << ((18446744073709551615) & 63))
        _ = v33
    }
    var v34 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
