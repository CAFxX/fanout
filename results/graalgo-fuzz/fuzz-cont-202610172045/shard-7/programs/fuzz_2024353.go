package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 int16, p1 bool, p2 uint64) uint8 {
    if (-636 >= 684) {
        var v6 uint64 = 882
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(p1, p0)
            var v8 bool = true
            if v8 {
                fmt.Println(v8)
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := []int{-2147483648, fi7}
                    v11 := GenS1{F0: 1, F1: (fi9 & (fi7 + int(p2)))}
                    switch "foo bar" {
                    case "a":
                        var v12 uint16 = 218
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                var v13 float64 = 1.5
                var v14 uint = 979
                _ = v13
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        var v15 uint32 = 179
        _ = v15
    }
    switch ("" + "") {
    case "hello":
        switch "foo bar" {
        case "a":
            if (!false) {
                v16 := GenS1{F0: 68, F1: -901}
                _ = v16
            }
        default:
            var v17 int32 = -859
            _ = v17
        }
    default:
        if (108 <= 42) {
            if (!p1) {
                for fi18 := 0; fi18 < 16; fi18++ {
                    p1 = (255 <= 407)
                    for fi19 := 0; fi19 < 2; fi19++ {
                        p2 = (p2 ^ 2)
                        v20 := []int{-1, (-(fi19))}
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        var v21 int64 = 1
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    var v22 int8 = 118
                    _ = fi18
                    _ = v22
                }
                var v23 float64 = 1.5
                var v24 int32 = -2147483648
                _ = v23
                _ = v24
            } else {
                p1 = (-1 >= 0)
            }
            var v25 string = ""
            v26 := GenS1{F0: 42, F1: 32768}
            _ = v25
            _ = v26
        } else {
            v27 := []int{2147483647, -9223372036854775808, -350}
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint8, p1 int16, p2 uint16) uint32 {
    v28 := GenS1{F0: -1, F1: 127}
    fmt.Println(p2)
    p1 = int16(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 65535
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 int = 0
    var v4 uint16 = 65535
    v5 := make(map[int]uint32)
    if (42 <= 120) {
        switch 1 {
        case 2147483647:
            v3 = (v3 & v3)
        default:
            v3 += (-(v3))
        }
        var v29 string = "a"
        if (!false) {
            if true {
                v1[v3] = v3
                v4 -= 42
                v3 += 0
            }
        } else {
            v30 := []int{(v3 - (v3 >> ((1) & 63)))}
            switch (v3 * v3) {
            case 283:
                v29 = v29
            default:
                var v31 float32 = -2.25
                _ = v31
            }
            _ = v30
        }
        _ = v29
    }
    v32 := GenS2{F0: ((^(v3)) <= 1), F1: -126, F2: 0}
    v3 -= v3
    switch (-(v3)) {
    case 115:
        delete(v1, 506)
    case -426:
        v33 := GenS2{F0: (true && true), F1: 100, F2: 42}
        _ = v33
    case 0:
        v3 -= 2147483647
    default:
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 uint = 255
            var v36 uint16 = v4
            if ((!("hello" >= "hello")) || false) {
                switch "hello" {
                case "":
                    v37 := GenS2{F0: (785 != 1), F1: 0, F2: 255}
                    _ = v37
                case "hello":
                    v4 = 153
                }
            }
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v4, v3)
    fmt.Println(v3, v4)
    var v38 int32 = -174
    v1[(v3 ^ 796)] = v3
    fmt.Println(v3, v3)
    v39 := GenS2{F0: false, F1: 127, F2: 221}
    switch (-(v3)) {
    case -1:
        var v40 bool = true
        _ = v40
    }
    if (42 > 42) {
        for fi41 := 0; fi41 < 17; fi41++ {
            v42 := []int{}
            switch 128 {
            case 537:
                var v43 int = v3
                _ = v43
            case 42:
                if len(v42) >= 2 {
                    v42 = v42[1:len(v42)]
                }
            case -9223372036854775808:
                var v44 float32 = 1.5
                _ = v44
            }
            _ = fi41
            _ = v42
        }
    }
    v45 := GenS3{F0: v4, F1: v38}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v5
}
