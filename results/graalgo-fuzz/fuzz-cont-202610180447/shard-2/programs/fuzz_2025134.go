package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 int = 32768
    v4 := make(map[uint32]int)
    var v5 uint = 0
    fmt.Println(v3, v1)
    v2["a"] = 65535
    if true {
        fmt.Println(v5)
        v1 = (!v1)
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint32 = 2147483648
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := []int{(v3 & int(v7))}
                v7 += 2147483648
                v10 := GenS1{F0: (v3 & (v3 & (fi6 * (fi8 | 0))))}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    } else {
        var v11 bool = (3.14159 <= ((((-2.25 * (0.1 / ((3.14159 - 1.5) / 1.0))) * 1e10) + ((3.14159 / 1e10) / 1.0)) + (3.14159 - 3.14159)))
        for fi12 := 0; fi12 < 5; fi12++ {
            v3 = (fi12 >> ((723) & 63))
            _ = fi12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint = (v5 / (((v5 * (v5 | 42))) | 1))
            if ((42 >= 4294967295) || (!v11)) {
                v5 -= v5
            }
            if (42 >= 853) {
                var v15 uint16 = 468
                _ = v15
            } else {
                var v16 int32 = -1
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        _ = v11
    }
    v5 = v5
    v4[0] = 42
    var v17 float32 = 1.5
    _ = len(v2)
    v18 := GenS1{F0: 4294967295}
    v17 += 1.5
    if v1 {
        v19 := GenS1{F0: (v3 ^ 32767)}
        if ((1e10 * (v17 + (float32(3.14159) / (v17 * 0.0)))) <= 3.14159) {
            v5 = v5
            var v20 float32 = ((3.14159 + 1e10) - -2.25)
            switch ("a" + "") {
            case "a":
                v21 := GenS1{F0: v3}
                _ = v21
            case "\n\t":
                fmt.Println(v1, v5)
            case "foo bar":
                v22 := GenS1{F0: v3}
                _ = v22
            }
            _ = v20
        } else {
            if false {
                v3 = v3
                var v23 uint64 = 164
                _ = v23
            } else {
                var v24 uint8 = 174
                _ = len(v4)
                for fi25 := 0; fi25 < 6; fi25++ {
                    fmt.Println(v24, v17)
                    var v26 int = v3
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        }
        var v27 int16 = 389
        _ = v19
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        if (v1 || v1) {
            var v29 uint16 = 883
            var v30 float64 = 1.5
            _ = v29
            _ = v30
        }
        fmt.Println(v3, fi28, v5)
        for fi31 := 0; fi31 < 12; fi31++ {
            v32 := GenS2{F0: -2147483648}
            var v33 int32 = -1
            var v34 int = (-(v3))
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi28
    }
    var v35 uint = (v5 - (v5 & v5))
    var v36 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
