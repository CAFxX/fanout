package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int8, p1 int, p2 uint) uint {
    fmt.Println(p1)
    switch (p1 / ((2147483647) | 1)) {
    case 128:
        p0 = p0
    case 32768:
        fmt.Println(p0, p2, p1)
    default:
        var v6 int64 = 0
        _ = v6
    }
    p0 = p0
    var v7 int32 = 2147483647
    if true {
        v8 := GenS1{F0: 42}
        switch p1 {
        case 786:
            switch (^(p1)) {
            case 9223372036854775807:
                var v9 uint64 = 229
                _ = v9
            }
        case 4294967295:
            var v10 int32 = v7
            _ = v10
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 18446744073709551615}
            var v13 uint16 = 1
            v14 := []int{p1}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v8
    } else {
        var v15 string = "foo bar"
        var v16 uint = p2
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 ^ 440)
}
func genFn1() int16 {
    v17 := GenS1{F0: 18446744073709551615}
    var v18 int = 127
    if (!true) {
        switch v18 {
        case -222:
            v18 += (v18 | v18)
        default:
            v18 += (v18 + v18)
        }
    } else {
        var v19 uint64 = 255
        fmt.Println(v19)
        _ = v19
    }
    v18 += v18
    if (false || (true && (1000 > -730))) {
        var v20 uint8 = 255
        _ = v20
    }
    _ = v17
    _ = v18
    return -32768
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 bool = (644 != -462)
    var v5 uint8 = 59
    var v21 int32 = 1
    switch 32768 {
    case -1:
        var v22 int8 = -40
        _ = v22
    case 285:
        v5 = 129
    }
    delete(v3, -439)
    for fi23 := 0; fi23 < 19; fi23++ {
        v5 += (^(v5))
        var v24 float64 = ((((0.0 - (-2.25 - -2.25)) * 0.1) * ((3.14159 * 0.1) + (1e10 - (1.5 * -2.25)))) / 3.14159)
        _ = fi23
        _ = v24
    }
    if true {
        v1 = (true && (65535 > 240))
    } else {
        if (42 == 0) {
            if true {
                var v25 uint16 = 255
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            v5 = 127
            var v27 int = (fi26 / ((32767) | 1))
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 string = ""
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS1{F0: 573}
    v1 = v4
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(v4, v5, v4)
        _ = fi31
    }
    var v32 float64 = float64(1.5)
    switch "foo bar" {
    case "\n\t":
        v33 := GenS2{F0: -9223372036854775808}
        _ = v33
    case "a":
        var v34 float64 = 1e10
        _ = v34
    case "foo bar":
        v21 = 42
    }
    var v35 uint = 278
    if true {
        delete(v3, 9223372036854775807)
    } else {
        for fi36 := 0; fi36 < 18; fi36++ {
            v21 = v21
            v37 := GenS2{F0: 0}
            _ = fi36
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        if (uint16(v35) == 0) {
            if (42 == 0) {
                for fi39 := 0; fi39 < 1; fi39++ {
                    var v40 int32 = v21
                    _ = fi39
                    _ = v40
                }
                switch ("foo bar" + ("foo bar" + "hello")) {
                case "":
                    switch (fi38 | (fi38 << (((v35 + 0)) & 63))) {
                    case 4294967295:
                        v41 := GenS2{F0: 939}
                        _ = v41
                    default:
                        v1 = v4
                    }
                case "hello":
                    for fi42 := 0; fi42 < 6; fi42++ {
                        v4 = (1 <= 65535)
                        _ = fi42
                    }
                default:
                    v43 := []int{fi38, (fi38 - fi38), (fi38 >> (((v35 >> ((genFn0(123, fi38, (v35 << (((v35 ^ v35)) & 63)))) & 63))) & 63))}
                    _ = v43
                }
            }
            var v44 float32 = (1.5 * 1.5)
            switch "foo bar" {
            case "\n\t":
                v45 := GenS2{F0: 42}
                _ = v45
            case "a":
                v46 := GenS3{F0: genFn0(104, fi38, genFn0(127, (fi38 / (((fi38 >> (((v35 >> (((^(v35))) & 63))) & 63))) | 1)), (^(v35))))}
                _ = v46
            default:
                var v47 int = (fi38 ^ -2147483648)
                _ = v47
            }
            _ = v44
        } else {
            v32 = (v32 + 0.1)
            v48 := GenS3{F0: genFn0(int8(fi38), fi38, (v35 + v35))}
            _ = v48
        }
        v5 = (v5 / ((2) | 1))
        _ = fi38
    }
    v1 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
