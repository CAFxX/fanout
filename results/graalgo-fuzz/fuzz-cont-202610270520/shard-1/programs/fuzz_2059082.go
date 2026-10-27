package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 int8, p2 uint) uint32 {
    if ((((0.0 - 1e10) + 0.1) <= (1e10 * (0.1 * 1.5))) || (!((p1 >> ((p2) & 63)) > (^(p1))))) {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                p1 = (^(p1))
                _ = fi8
            }
            _ = fi7
        }
        p0 += 873
    } else {
        var v9 string = "foo bar"
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 1, F1: 482, F2: 86}
            _ = fi10
            _ = v11
        }
        fmt.Println(p0, p2, v9)
        _ = v9
    }
    v12 := GenS1{F0: 2, F1: 459, F2: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2
}

func main() {
    var v1 int16 = 0
    var v2 int = 754
    var v3 int8 = -39
    v4 := make(map[uint32]string)
    var v5 uint = 26
    var v6 uint16 = 0
    v13 := GenS1{F0: uint8(v5), F1: 0, F2: 0}
    _ = len(v4)
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            if (4294967295 >= 4) {
                for fi15 := 0; fi15 < 10; fi15++ {
                    fmt.Println(v2, fi15)
                    v4[2147483648] = "hello"
                    _ = fi15
                }
            }
            if (true || false) {
                v16 := GenS1{F0: 105, F1: 345, F2: 660}
                v5 += v5
                _ = v16
            }
            v17 := GenS1{F0: 2, F1: 0, F2: 618}
            _ = fi14
            _ = v17
        }
        v5 -= (^(v5))
    }
    v18 := GenS2{F0: 4294967295}
    switch "foo bar" {
    case "\n\t":
        _ = len(v4)
    case "":
        if false {
            v19 := GenS1{F0: 0, F1: (v6 - 65535), F2: 2}
            _ = v19
        } else {
            var v20 string = "\n\t"
            v1 += 250
            _ = v20
        }
    default:
        fmt.Println(v3, v1)
    }
    switch "\n\t" {
    case "hello":
        var v21 uint = (v5 >> (((v5 | (v5 | v5))) & 63))
        _ = v21
    case "\n\t":
        v22 := GenS1{F0: 46, F1: 42, F2: 0}
        _ = v22
    default:
        v23 := []int{(v2 >> ((v5) & 63))}
        _ = v23
    }
    var v24 bool = (false || true)
    if v24 {
        fmt.Println(v1, v3)
    } else {
        v6 += v6
        var v25 float32 = (float32((0.0 - -2.25)) / 3.14159)
        switch "foo bar" {
        case "a":
            v6 += (v6 / ((v6) | 1))
        case "":
            var v26 int = v2
            _ = v26
        }
        _ = v25
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        delete(v4, 2147483648)
        for fi28 := 0; fi28 < 10; fi28++ {
            v2 += (fi28 << ((65535) & 63))
            switch (((("a" + (("a" + "hello") + "\n\t")) + (("hello" + "\n\t") + "hello")) + "foo bar") + "") {
            case "foo bar":
                v29 := GenS1{F0: uint8(v5), F1: (v6 | (^(v6))), F2: 42}
                _ = v29
            case "":
                fmt.Println(fi28, v3)
            }
            fmt.Println(v6, fi28)
            _ = fi28
        }
        _ = fi27
    }
    var v30 uint8 = 42
    var v31 int64 = -1
    switch (^(v2)) {
    case 32768:
        switch -1 {
        case 32767:
            var v32 uint = (v5 / ((v5) | 1))
            _ = v32
        case -348:
            v33 := GenS3{F0: v31}
            _ = v33
        default:
            if v24 {
                delete(v4, uint32(v6))
                if v24 {
                    fmt.Println(v31, v24)
                }
            } else {
                v34 := GenS1{F0: v30, F1: (v6 % ((uint16(v5)) | 1)), F2: genFn0(358, int8(v1), 2)}
                var v35 uint16 = (v6 << ((v5) & 63))
                _ = v34
                _ = v35
            }
        }
    case -1:
        delete(v4, 7)
    default:
        v36 := []int{v2, -461, (v2 << (((v5 ^ (v5 | 512))) & 63))}
        _ = v36
    }
    v1 -= (v1 * v1)
    var v37 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
