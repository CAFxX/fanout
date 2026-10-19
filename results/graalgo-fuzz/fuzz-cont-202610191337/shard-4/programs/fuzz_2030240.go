package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0() int16 {
    v7 := []int{855, -7, 42, 128}
    if ((false || true) || false) {
        var v8 int = -2147483648
        if (85 >= 40) {
            fmt.Println(v8)
            v8 -= v8
            if ((false && (42 >= 0)) && (684 > 252)) {
                var v9 uint32 = 0
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := GenS1{F0: 2, F1: 127}
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
        } else {
            v8 = v8
            if len(v7) > 0 {
                v7[0] = v8
            }
            v8 = v8
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v8
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint32 = 42
            v14 := GenS2{F0: (v13 * 2147483647)}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: 758}
        _ = v15
    }
    _ = len(v7)
    _ = v7
    return -32768
}

func main() {
    v1 := []int{-1, 32768, 2147483647}
    var v2 int64 = 42
    var v3 bool = false
    var v4 string = ("hello" + ("hello" + "foo bar"))
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int8 = 127
        _ = fi16
        _ = v17
    }
    switch v4 {
    case "a":
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        if false {
            v19 := GenS2{F0: 65535}
            _ = v19
        }
        var v20 bool = true
        v21 := GenS3{F0: fi18, F1: int16(fi18), F2: v20}
        _ = fi18
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 981}
    delete(v5, 886)
    if v3 {
        var v23 uint16 = 195
        _ = v23
    } else {
        var v24 uint = 4294967295
        _ = v24
    }
    fmt.Println(v4, v4, v2)
    if (1 < v2) {
        if (!v3) {
            v4 = ("" + (v4 + "a"))
        }
        v3 = (v2 >= -9223372036854775808)
    } else {
        var v25 uint8 = 255
        _ = v25
    }
    switch 42 {
    case 32767:
        if v3 {
            v4 = v4
        } else {
            v3 = v3
            for fi26 := 0; fi26 < 10; fi26++ {
                v4 = v4
                var v27 uint64 = 42
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 1; fi28++ {
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 int = fi28
                    fmt.Println(fi28, v2)
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        v2 -= v2
                    case "hello":
                        v31 := GenS4{F0: 2, F1: 65535}
                        _ = v31
                    }
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v2)
                var v32 int8 = 1
                _ = fi28
                _ = v32
            }
        }
    case -1:
        fmt.Println(v3, v2, v3)
    case 32768:
        for fi33 := 0; fi33 < 4; fi33++ {
            for fi34 := 0; fi34 < 19; fi34++ {
                if ((v2 ^ 485) > (v2 * (-(v2)))) {
                    var v35 float64 = -2.25
                    _ = v35
                }
                if (1 == 4294967295) {
                    fmt.Println(v2, v3, fi34)
                } else {
                    var v36 uint32 = 2147483648
                    v3 = (v3 || (!v3))
                    _ = v36
                }
                v3 = (2147483647 < (fi33 & 9223372036854775807))
                _ = fi34
            }
            _ = fi33
        }
    default:
        switch v4 {
        case "a":
            switch 127 {
            case 694:
                v37 := GenS3{F0: -9223372036854775808, F1: genFn0(), F2: false}
                _ = v37
            default:
                v38 := GenS2{F0: 1}
                _ = v38
            }
        }
    }
    var v39 uint8 = 1
    v4 = v4
    v1 = append(v1, 32768)
    var v40 uint32 = 42
    for fi41 := 0; fi41 < 20; fi41++ {
        if v3 {
            var v42 bool = true
            var v43 int64 = (v2 * v2)
            _ = v42
            _ = v43
        }
        _ = fi41
    }
    var v44 uint32 = 65535
    v45 := v5[v44]
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
