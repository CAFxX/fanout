package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint32) uint16 {
    var v4 int16 = -32768
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{42, -1, fi5}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 153
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = float32((-2.25 / (((3.14159 / 1e10) * float64(1e10)) + 0.0)))
    var v3 float32 = 1e10
    v2 = v3
    var v7 uint = 42
    fmt.Println(v7, v2, v2)
    var v8 uint32 = 1
    v2 = v2
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int32 = -219
        switch 4294967295 {
        case 32767:
            if false {
                if (4294967295 >= (fi9 ^ 42)) {
                    _ = len(v1)
                }
            }
        case 261:
            v8 += (v8 >> (((^(v7))) & 63))
        default:
            v11 := GenS1{F0: ("foo bar" + (("foo bar" + ("a" + "")) + ""))}
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v12 := v1["hello"]
    var v13 int64 = -1
    if (v12 || (255 <= 1)) {
        var v14 int16 = 32767
        var v15 int64 = -921
        _ = v14
        _ = v15
    }
    if (uint16(v7) >= 0) {
        if v12 {
            var v16 int8 = -128
            _ = v16
        } else {
            if (((v8 + (v8 | 2147483648)) == v8) || (903 <= -1)) {
                v1["\n\t"] = (!((v12 && v12) && (v3 == 1e10)))
                v2 = v3
            } else {
                fmt.Println(v12)
                fmt.Println(v8, v8)
                var v17 int8 = 0
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 string = "a"
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v13, v7, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
