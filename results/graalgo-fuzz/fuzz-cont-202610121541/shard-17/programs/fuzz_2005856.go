package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 int32, p1 bool) uint8 {
    var v7 float64 = 3.14159
    var v8 uint32 = 0
    var v9 float32 = float32(-2.25)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 bool, p1 int64, p2 string) int8 {
    fmt.Println(p0, p1, p2)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    v1 := []uint32{961, 2147483648, 1, 4294967295}
    v2 := make(map[uint32]uint32)
    v3 := []int{11}
    var v4 uint32 = 2147483648
    var v5 uint = 0
    var v6 uint16 = 32767
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint32 = (v4 - (v4 << ((v5) & 63)))
            _ = len(v1)
            _ = len(v1)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if false {
        if (-436 >= 42) {
            fmt.Println(v6, v4)
        } else {
            if (true && false) {
                switch "\n\t" {
                case "a":
                    _ = len(v1)
                case "":
                    if (!true) {
                        var v13 bool = true
                        _ = v13
                    }
                case "hello":
                    var v14 uint64 = 4294967295
                    _ = v14
                }
                var v15 uint8 = 236
                v16 := GenS1{F0: v15}
                _ = v15
                _ = v16
            }
            v3 = append(v3, 32767)
        }
    } else {
        v6 -= 255
        switch 127 {
        case -1:
            v17 := GenS1{F0: 0}
            _ = v17
        }
        v2[669] = v4
    }
    if false {
        v18 := GenS1{F0: 1}
        _ = v18
    }
    var v19 uint32 = v4
    var v20 uint32 = (v19 / ((v19) | 1))
    switch "" {
    case "\n\t":
        _ = len(v2)
    }
    switch "hello" {
    case "hello":
        fmt.Println(v19, v6)
    case "":
        var v21 int = -359
        _ = v21
    default:
        if (0.0 <= 1.5) {
            v6 += (^(v6))
            var v22 uint32 = v4
            var v23 string = ""
            _ = v22
            _ = v23
        }
    }
    var v24 uint32 = v20
    if false {
        switch "" {
        case "hello":
            v5 = 791
        case "":
            v19 += v4
        case "a":
            var v25 uint32 = (v24 << ((42) & 63))
            _ = v25
        }
        v1 = append(v1, v4)
    }
    fmt.Println(v19, v20, v4)
    fmt.Println(v19, v24)
    for fi26 := 0; fi26 < 9; fi26++ {
        switch 1 {
        case 127:
            var v27 uint = (v5 % ((v5) | 1))
            _ = v27
        case 179:
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := GenS1{F0: 222}
                _ = fi28
                _ = v29
            }
        case 1:
            fmt.Println(v24, v20, v20)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v30 int32 = 0
        _ = fi26
        _ = v30
    }
    if true {
        fmt.Println(v4)
        if true {
            var v31 uint32 = v24
            v4 = v20
            fmt.Println(v5)
            _ = v31
        } else {
            v32 := GenS1{F0: 158}
            v33 := GenS1{F0: 42}
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 uint8 = 127
                var v36 uint32 = uint32(v35)
                for fi37 := 0; fi37 < 20; fi37++ {
                    v35 += (^(v35))
                    switch "a" {
                    case "foo bar":
                        v38 := GenS1{F0: v35}
                        _ = v38
                    case "hello":
                        var v39 int = (fi37 << ((4294967295) & 63))
                        _ = v39
                    }
                    _ = fi37
                }
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v32
            _ = v33
        }
        v4 -= v20
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            if len(v3) > 0 {
                v3[0] = fi40
            }
            v6 = (v6 / ((v6) | 1))
            var v41 uint = (^(v5))
            _ = fi40
            _ = v41
        }
        if (!(1e10 != ((float32(0.1) + 0.0) - (3.14159 - 1e10)))) {
            fmt.Println(v6, v24, v6)
        } else {
            fmt.Println(v19)
        }
    }
    if (0 < 42) {
        for fi42 := 0; fi42 < 5; fi42++ {
            if false {
                var v43 int8 = 42
                _ = v43
            }
            if len(v1) > 0 {
                v1[0] = (v24 | (^(v4)))
            }
            _ = fi42
        }
        var v44 int16 = 32767
        _ = v44
    } else {
        v5 = v5
    }
    var v45 int = 4294967295
    fmt.Println(v19, v45)
    switch (v45 ^ -2147483648) {
    case 0:
        switch v45 {
        case 1:
            if (((float32(1e10) == -2.25) || true) || (!false)) {
                v5 = (v5 + (v5 - uint(v6)))
                v45 -= 442
            } else {
                v45 += (^(v45))
                for fi46 := 0; fi46 < 18; fi46++ {
                    v47 := GenS1{F0: 127}
                    _ = fi46
                    _ = v47
                }
            }
        case 32768:
            for fi48 := 0; fi48 < 7; fi48++ {
                v49 := v2[v19]
                if (v5 == (v5 + (v5 / (((v5 ^ 141)) | 1)))) {
                    v2[v19] = v19
                    v50 := GenS2{F0: v5, F1: genFn1(false, int64(v45), ""), F2: -897}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v50
                } else {
                    var v51 uint = v5
                    _ = v51
                }
                _ = fi48
                _ = v49
            }
        case -9223372036854775808:
            v24 = (^(v24))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
