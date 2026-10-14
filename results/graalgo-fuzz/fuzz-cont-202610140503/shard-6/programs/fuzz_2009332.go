package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint64; F2 string }
type GenS4 struct { F0 int }
func genFn0(p0 bool, p1 int64, p2 uint16) int16 {
    fmt.Println(p1)
    for fi7 := 0; fi7 < 7; fi7++ {
        p2 = (p2 / (((^(p2))) | 1))
        v8 := GenS1{F0: 42, F1: fi7}
        v9 := GenS1{F0: uint(p2), F1: -940}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 1000
}
func genFn1(p0 bool, p1 string, p2 int) uint16 {
    for fi10 := 0; fi10 < 13; fi10++ {
        switch (fi10 - (p2 + fi10)) {
        case 1:
            if p0 {
                p2 -= 9223372036854775807
                v11 := GenS1{F0: 42, F1: (-(p2))}
                v12 := GenS2{F0: 0}
                _ = v11
                _ = v12
            }
        case 591:
            p2 = (fi10 >> ((255) & 63))
        }
        if p0 {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS3{F0: -1, F1: 484, F2: "a"}
                var v15 float64 = ((-2.25 - (1.5 / -2.25)) / 1.5)
                p0 = (100 <= 100)
                fmt.Println(v15, v15)
                if p0 {
                    fmt.Println(p1, v15, p0)
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            fmt.Println(p1)
            p1 = p1
            switch (fi10 >> ((978) & 63)) {
            case 1:
                p1 = (p1 + p1)
            case -264:
                if p0 {
                    var v16 int32 = 866
                    v17 := []int{(p2 ^ fi10), p2}
                    var v18 int64 = 1
                    _ = v16
                    _ = v17
                    _ = v18
                }
            case 9223372036854775807:
                v19 := []int{2147483647, p2}
                _ = v19
            }
        }
        _ = fi10
    }
    switch (p2 / ((0) | 1)) {
    case 1:
        v20 := GenS2{F0: 42}
        _ = v20
    case 42:
        for fi21 := 0; fi21 < 9; fi21++ {
            if (-1 < -32768) {
                switch ((p1 + p1) + p1) {
                case "":
                    var v22 int32 = 0
                    _ = v22
                default:
                    v23 := []int{}
                    _ = v23
                }
                var v24 uint32 = 340
                var v25 float32 = 0.0
                _ = v24
                _ = v25
            }
            p0 = (!p0)
            _ = fi21
        }
    default:
        var v26 uint16 = 255
        _ = v26
    }
    switch 9223372036854775807 {
    case 2147483647:
        fmt.Println(p2)
    case -162:
        v27 := []int{p2}
        _ = v27
    default:
        p1 = p1
    }
    var v28 float32 = (((1.5 / (-2.25 / 0.1)) / 3.14159) - 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 65535
}

func main() {
    var v1 uint = 147
    var v2 bool = (-241 != 4294967295)
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    var v5 string = "hello"
    var v6 uint8 = 42
    var v29 int8 = 1
    fmt.Println(v6, v5, v3)
    var v30 uint8 = (v6 & 132)
    v3 -= v3
    for fi31 := 0; fi31 < 5; fi31++ {
        _ = len(v4)
        for fi32 := 0; fi32 < 20; fi32++ {
            v30 -= (v6 >> ((uint(v3)) & 63))
            if v2 {
                v33 := GenS4{F0: fi32}
                _ = v33
            }
            var v34 int32 = 42
            fmt.Println(v5, v34, v2)
            _ = fi32
            _ = v34
        }
        _ = fi31
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 float32 = -2.25
            if ((!false) || (77 < 42)) {
                if v2 {
                    fmt.Println(v6, v2)
                    delete(v4, (v3 | (v3 << ((4294967295) & 63))))
                } else {
                    v38 := v4[v3]
                    v39 := GenS1{F0: (v1 | (v1 << (((v1 / (((v1 | (v1 - 18446744073709551615))) | 1))) & 63))), F1: fi36}
                    _ = v38
                    _ = v39
                }
            } else {
                v40 := GenS1{F0: (v1 - uint(v6)), F1: (fi36 << ((uint(v30)) & 63))}
                _ = v40
            }
            var v41 bool = (1e10 < 1e10)
            _ = fi36
            _ = v37
            _ = v41
        }
        _ = fi35
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        _ = len(v4)
        _ = fi42
    }
    v4[(v3 - (v3 >> (((v1 & v1)) & 63)))] = v3
    var v43 bool = v2
    _ = len(v4)
    var v44 float32 = (float32(1e10) - (0.0 - (((3.14159 - 1.5) - (3.14159 * 1e10)) * -2.25)))
    v44 = v44
    v43 = false
    v30 = v30
    v44 = v44
    switch ("a" + (v5 + v5)) {
    case "a":
        if (!((v43 || v43) || v2)) {
            v1 = 657
        }
    case "":
        var v45 bool = (1 >= 2)
        _ = v45
    case "foo bar":
        if v43 {
            var v46 bool = (v2 && v43)
            if v2 {
                v29 = v29
            }
            _ = v46
        } else {
            fmt.Println(v5, v30)
        }
    }
    switch -1 {
    case -9223372036854775808:
        switch int(v3) {
        case -2147483648:
            v47 := GenS2{F0: genFn1(v2, (v5 + "\n\t"), 2147483647)}
            _ = v47
        }
    default:
        var v48 uint16 = genFn1(v2, v5, -471)
        _ = v48
    }
    v49 := []int{573}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v49
    _ = v5
    _ = v6
}
