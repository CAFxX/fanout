package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int32; F1 int32; F2 int8 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint, p1 int, p2 bool) int64 {
    v7 := []int{p1, -5, (p1 >> ((p0) & 63)), (p1 & (-(p1)))}
    for fi8 := 0; fi8 < 20; fi8++ {
        if (148 <= 255) {
            if false {
                p2 = (909 > 1000)
                p2 = false
                v9 := GenS1{F0: (fi8 / ((p1) | 1))}
                _ = v9
            } else {
                v10 := GenS2{F0: 0, F1: 208, F2: 18446744073709551615}
                var v11 uint = p0
                _ = v10
                _ = v11
            }
        }
        _ = fi8
    }
    if (!false) {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS2{F0: 2, F1: 214, F2: 82}
            fmt.Println(p1, fi12, fi12)
            _ = fi12
            _ = v13
        }
    } else {
        fmt.Println(p1, p2, p1)
        for fi14 := 0; fi14 < 14; fi14++ {
            switch (fi14 - p1) {
            case 0:
                var v15 uint16 = 1
                _ = v15
            }
            var v16 uint64 = 42
            fmt.Println(p2)
            _ = fi14
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            if (p2 || (p2 && false)) {
                v18 := GenS1{F0: (p1 + -637)}
                var v19 uint = 42
                fmt.Println(fi17)
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    p2 = ((float32((((1e10 + 1.5) * -2.25) / 1.0)) - (3.14159 / 1e10)) <= float32(1.5))
                    p1 -= fi20
                    fmt.Println(p0, fi17)
                    _ = fi20
                }
            }
            fmt.Println(p1)
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 string = "\n\t"
    var v2 string = "hello"
    v3 := make(map[uint32]string)
    v4 := []uint32{2147483648}
    v5 := make(map[uint32]int)
    var v6 int32 = -1
    v6 += (v6 + v6)
    if true {
        var v21 int32 = (^(v6))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ((0 <= -9223372036854775808) || true) {
            v21 = v21
            fmt.Println(v6, v2, v2)
            v22 := GenS3{F0: -787, F1: genFn0(42, -2147483648, false)}
            _ = v22
        } else {
            if (!(v6 < -113)) {
                var v23 uint16 = 42
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v23
            }
            v24 := GenS4{F0: 9223372036854775807, F1: 1, F2: 1}
            _ = v24
        }
        _ = v21
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS3{F0: (fi25 | fi25), F1: 42}
        v27 := GenS5{F0: (v6 % ((v6) | 1)), F1: (v6 >> ((0) & 63)), F2: 42}
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint16 = 32767
    var v29 int32 = v6
    fmt.Println(v28, v6)
    for fi30 := 0; fi30 < 17; fi30++ {
        if false {
            var v31 int32 = 0
            _ = v31
        } else {
            switch (^(fi30)) {
            case -177:
                var v32 bool = (!(18446744073709551615 < 42))
                _ = v32
            case 128:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v29 -= (v6 + 50)
            fmt.Println(v2, fi30, v2)
        }
        if true {
            delete(v5, 2147483647)
            var v33 uint64 = 0
            _ = len(v4)
            _ = v33
        } else {
            var v34 int8 = -128
            _ = v34
        }
        _ = fi30
    }
    v6 = 42
    v35 := GenS6{F0: true, F1: 100}
    v6 -= (v29 - int32(v28))
    v29 = (v29 >> ((1) & 63))
    v36 := GenS7{F0: -32768}
    v1 = (("foo bar" + "foo bar") + v2)
    var v37 int32 = v29
    var v38 int8 = -128
    fmt.Println(v38, v1)
    fmt.Println(v29, v1, v28)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
