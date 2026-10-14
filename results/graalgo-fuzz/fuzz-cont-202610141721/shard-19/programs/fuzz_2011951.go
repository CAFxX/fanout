package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 int) int32 {
    switch ("hello" + "foo bar") {
    case "a":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p0)
        case "a":
            fmt.Println(p0)
        }
    case "hello":
        v5 := GenS1{F0: -422, F1: 0}
        _ = v5
    case "\n\t":
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int32 = -35
            v8 := []int{-2147483648, (fi6 / ((p0) | 1)), (p0 | (-(fi6)))}
            v9 := GenS1{F0: -1, F1: 42}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    var v10 uint = 65535
    v11 := []int{-9223372036854775808}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 bool = (false && (!true))
        v14 := GenS2{F0: (!false), F1: 437, F2: false}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int32 = -2147483648
    _ = p0
    _ = v10
    _ = v11
    _ = v15
    return 531
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 728
    var v3 float64 = 1.5
    var v4 int64 = 23
    v16 := []int{-1, 128, 414}
    v16 = append(v16, 1)
    _ = len(v16)
    switch -378 {
    case 32767:
        var v17 string = v1
        _ = v17
    case -194:
        var v18 string = v1
        _ = v18
    case 127:
        fmt.Println(v2, v3, v3)
    default:
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 float64 = (v3 - (1.5 / ((v3 / float64(-2.25)) - v3)))
            v4 += -677
            _ = fi19
            _ = v20
        }
    }
    var v21 int = 0
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        if true {
            v23 := GenS1{F0: int64(v21), F1: 156}
            v2 = (v2 ^ v2)
            v24 := GenS3{F0: v1, F1: (true && false)}
            _ = v23
            _ = v24
        }
        _ = len(v16)
        _ = fi22
    }
    if false {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            if len(v16) > 0 {
                v16[0] = v21
            }
            var v26 int = (v21 << ((4294967295) & 63))
            var v27 float32 = (3.14159 + float32((float64(0.1) / (v3 * float64(float32(float64(0.0)))))))
            _ = fi25
            _ = v26
            _ = v27
        }
        if ((v2 & v2) <= 2147483648) {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            for fi28 := 0; fi28 < 16; fi28++ {
                if (true || false) {
                    v16 = append(v16, int(v2))
                    var v29 int16 = -32768
                    _ = v29
                } else {
                    var v30 int32 = -594
                    var v31 bool = (v30 != v30)
                    _ = v30
                    _ = v31
                }
                var v32 float32 = 0.1
                _ = fi28
                _ = v32
            }
        }
    } else {
        if len(v16) > 0 {
            v16[0] = 42
        }
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 int32 = -950
        switch v21 {
        case -368:
            v21 += (fi33 / (((fi33 >> ((192) & 63))) | 1))
        case 128:
            v35 := GenS4{F0: v1, F1: v34, F2: 1000}
            _ = v35
        case 4294967295:
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint16 = 2
                v1 = (v1 + v1)
                _ = fi36
                _ = v37
            }
        }
        _ = fi33
        _ = v34
    }
    v3 -= (3.14159 - v3)
    v3 -= v3
    switch (v21 % ((1) | 1)) {
    case -254:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    case -111:
        v1 = v1
    case 824:
        if len(v16) > 0 {
            v16[0] = 32767
        }
    default:
        for fi38 := 0; fi38 < 5; fi38++ {
            switch (v21 & v21) {
            case 0:
                v2 = (^(v2))
            case -841:
                var v39 int = (fi38 | fi38)
                _ = v39
            case 127:
                v40 := GenS1{F0: (-(v4)), F1: 2}
                _ = v40
            }
            _ = fi38
        }
    }
    fmt.Println(v4, v2, v1)
    fmt.Println(v3, v2, v21)
    v2 = 2147483648
    v1 = "\n\t"
    v1 = v1
    v41 := GenS3{F0: v1, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
}
