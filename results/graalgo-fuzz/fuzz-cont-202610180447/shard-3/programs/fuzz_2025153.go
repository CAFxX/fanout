package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string, p1 int64) uint16 {
    if (!true) {
        v5 := GenS1{F0: p1}
        v6 := GenS1{F0: (-(p1))}
        p0 = ""
        _ = v5
        _ = v6
    }
    p0 = p0
    if (!true) {
        var v7 float64 = (0.0 + 0.0)
        v8 := []int{127}
        _ = v7
        _ = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 string = "a"
    var v4 int16 = 584
    v3 = v3
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 string = v3
        v12 := GenS1{F0: 917}
        var v13 int8 = 0
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 int64 = -273
    var v15 uint = 2
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: (v14 & (v14 - v14))}
        fmt.Println(v14, v3)
        _ = fi16
        _ = v17
    }
    v18 := v2[(v3 + "foo bar")]
    v19 := GenS1{F0: v14}
    var v20 uint64 = 699
    var v21 bool = (-14 == 127)
    v18 += (v18 - (v18 % ((v18) | 1)))
    if true {
        fmt.Println(v3, v3)
        var v22 float32 = 1.5
        for fi23 := 0; fi23 < 16; fi23++ {
            v21 = ((v18 % ((fi23) | 1)) != fi23)
            var v24 uint8 = 2
            for fi25 := 0; fi25 < 14; fi25++ {
                v14 -= v14
                var v26 string = ("" + (((v3 + (v3 + v3)) + v3) + (v3 + v3)))
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        if true {
            var v27 int64 = 0
            _ = v27
        }
        v28 := GenS1{F0: (v14 * (v14 << (((v15 & v15)) & 63)))}
        v29 := []int{(v18 - 0), (v18 ^ 42), v18, 4294967295}
        _ = v28
        _ = v29
    }
    if (777 == 2147483648) {
        v20 += (v20 >> (((^(v15))) & 63))
        var v30 int = (v18 + v18)
        for fi31 := 0; fi31 < 19; fi31++ {
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 int32 = -1
                _ = fi32
                _ = v33
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                v14 = int64(v20)
                _ = fi34
            }
            switch v18 {
            case 885:
                v35 := GenS1{F0: (v14 % ((9223372036854775807) | 1))}
                _ = v35
            case 95:
                switch v3 {
                case "hello":
                    v36 := GenS1{F0: (v14 - 1)}
                    _ = v36
                case "\n\t":
                    if true {
                        var v37 uint = v15
                        v14 = 0
                        _ = v37
                    }
                }
            case -57:
                v14 += (v14 * (^(v14)))
            }
            _ = fi31
        }
        _ = v30
    } else {
        v38 := GenS1{F0: 1}
        v39 := []int{v18, 32768, 0}
        _ = v38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
