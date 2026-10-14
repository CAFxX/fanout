package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int32; F1 bool; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 uint64; F2 bool }
func genFn0(p0 uint, p1 uint16, p2 uint) uint64 {
    if (2147483647 >= 2147483647) {
        p2 -= 65535
        switch 1 {
        case -9223372036854775808:
            for fi5 := 0; fi5 < 10; fi5++ {
                var v6 int32 = 647
                var v7 uint64 = 42
                v8 := GenS1{F0: -128}
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
        case 32768:
            v9 := []int{0, 127, -9223372036854775808}
            _ = v9
        default:
            p0 = p2
        }
        if (!(1 > 651)) {
            var v10 float64 = 1e10
            if true {
                p2 = (p2 - 762)
                p2 = (p2 - uint(p1))
                var v11 float32 = 0.0
                _ = v11
            }
            _ = v10
        }
    }
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 string = ("hello" + "hello")
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    var v4 float32 = 1.5
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            if (4294967295 != 204) {
                v14 := GenS2{F0: 874}
                v1 = v1
                _ = v14
            }
            _ = fi13
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v4, fi15, v4)
        var v16 int8 = 1
        v1 = v1
        _ = fi15
        _ = v16
    }
    v17 := v2[0]
    if true {
        for fi18 := 0; fi18 < 4; fi18++ {
            v4 = v4
            _ = len(v2)
            fmt.Println(v17)
            _ = fi18
        }
    }
    if (!false) {
        var v19 uint8 = 127
        v19 += v19
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 5; fi21++ {
                if true {
                    v19 += 0
                    if false {
                        v22 := []int{(fi20 % (((fi21 * fi21)) | 1))}
                        _ = v22
                    } else {
                        if true {
                            delete(v2, 2)
                            v1 = v1
                            var v23 uint = 1
                            _ = v23
                        } else {
                            v4 = 3.14159
                        }
                        switch ("a" + v1) {
                        case "\n\t":
                            v24 := GenS3{F0: 65535, F1: 172}
                            _ = v24
                        }
                    }
                    switch -9223372036854775808 {
                    case 857:
                        v19 -= v19
                    }
                } else {
                    switch fi21 {
                    case -362:
                        v25 := GenS4{F0: -2147483648, F1: false, F2: 196}
                        _ = v25
                    case 252:
                        v17 = v17
                    case 9223372036854775807:
                        delete(v3, fi21)
                    }
                    v26 := GenS5{F0: fi21}
                    var v27 float32 = (-2.25 / (v4 * v4))
                    _ = v26
                    _ = v27
                }
                if ((!(!(!true))) || true) {
                    var v28 float64 = -2.25
                    _ = v28
                }
                _ = fi21
            }
            v19 -= v19
            switch v17 {
            case "":
                if (false && (true || (!(false || (870 != 1))))) {
                    v4 = -2.25
                    for fi29 := 0; fi29 < 12; fi29++ {
                        if true {
                            fmt.Println(v17, v19)
                            var v30 int = -635
                            v31 := GenS6{F0: (v30 | (v30 | fi20))}
                            _ = v30
                            _ = v31
                        }
                        _ = fi29
                    }
                }
            case "foo bar":
                var v32 uint = 0
                _ = v32
            case "a":
                for fi33 := 0; fi33 < 12; fi33++ {
                    v34 := []int{}
                    var v35 float64 = 1.5
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            _ = fi20
        }
        _ = v19
    } else {
        if true {
            switch 0 {
            case 9223372036854775807:
                var v36 int64 = 1
                _ = v36
            case 32768:
                v17 = ("a" + v17)
            case 32767:
                v37 := GenS3{F0: 875, F1: 18446744073709551615}
                _ = v37
            }
        }
        if true {
            var v38 int = 128
            _ = v38
        } else {
            var v39 int8 = -1
            _ = v39
        }
    }
    fmt.Println(v17, v4, v17)
    v4 += ((v4 * 1e10) + 1e10)
    if true {
        v17 = v1
        v40 := GenS7{F0: 95, F1: 42, F2: true}
        _ = v40
    }
    switch 127 {
    case -214:
        var v41 int64 = 0
        _ = v41
    case -893:
        v4 += (1.5 / v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
