package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int32, p1 uint16) uint8 {
    var v7 int = 127
    v8 := []int{(v7 | v7)}
    v9 := GenS1{F0: 18446744073709551615, F1: 84, F2: false}
    p1 = 255
    v7 = v7
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 6
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = ((0 <= -943) || false)
    v3 := []int{290, 42, 42, 4294967295}
    var v4 int = -273
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    for fi10 := 0; fi10 < 7; fi10++ {
        if ((true && v2) && (42 <= 100)) {
            v11 := GenS2{F0: 255}
            v12 := GenS1{F0: 18446744073709551615, F1: 32767, F2: (!(459 > 18446744073709551615))}
            v2 = v2
            _ = v11
            _ = v12
        }
        v2 = v2
        _ = fi10
    }
    var v13 float64 = 1e10
    var v14 int8 = 0
    if (v2 || false) {
        var v15 string = ("foo bar" + "a")
        if (!v2) {
            switch (v4 & v4) {
            case 2147483647:
                _ = len(v5)
            case -349:
                _ = len(v6)
            default:
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 int16 = 1000
                    v18 := GenS2{F0: 42}
                    v2 = (1 >= 255)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            var v19 int = (v4 << ((925) & 63))
            if ((!(!true)) && (v2 || (true && ((1.5 - 1e10) > 0.1)))) {
                v15 = ((v15 + v15) + "a")
            }
            _ = v19
        }
        _ = v15
    } else {
        var v20 float32 = 0.1
        switch "" {
        case "foo bar":
            if v2 {
                fmt.Println(v13, v14)
                for fi21 := 0; fi21 < 17; fi21++ {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v23 := GenS3{F0: v4, F1: 0}
                        v13 = (3.14159 / 1.5)
                        _ = fi22
                        _ = v23
                    }
                    v6[v4] = ("a" + "\n\t")
                    for fi24 := 0; fi24 < 3; fi24++ {
                        fmt.Println(v20)
                        _ = fi24
                    }
                    _ = fi21
                }
            }
        case "\n\t":
            for fi25 := 0; fi25 < 11; fi25++ {
                v4 -= v4
                _ = fi25
            }
        default:
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint32 = 65535
                _ = fi26
                _ = v27
            }
        }
        _ = v20
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS1{F0: 42, F1: 631, F2: (644 == 18446744073709551615)}
        _ = fi28
        _ = v29
    }
    v30 := GenS4{F0: 511}
    if v2 {
        fmt.Println(v14, v13)
        var v31 uint16 = 22
        v32 := GenS1{F0: 255, F1: 1, F2: ((1.5 / v13) >= v13)}
        _ = v31
        _ = v32
    } else {
        v33 := GenS2{F0: 189}
        _ = v33
    }
    var v34 uint8 = genFn0(2147483647, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
