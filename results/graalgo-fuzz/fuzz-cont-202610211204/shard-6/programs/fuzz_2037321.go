package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32, p1 uint) uint8 {
    var v5 float64 = (3.14159 + (((0.0 / 1.0) + float64(float32((1e10 / 1e10)))) * 3.14159))
    switch 4294967295 {
    case 4294967295:
        var v6 uint64 = 0
        _ = v6
    case 42:
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{(fi7 | (fi7 / ((fi7) | 1))), fi7}
            var v9 string = ""
            _ = fi7
            _ = v8
            _ = v9
        }
    case 523:
        v10 := GenS1{F0: 1}
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{}
        switch fi11 {
        case 127:
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 float32 = (0.0 + 3.14159)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi13
                _ = v14
            }
        default:
            if false {
                p1 = 65535
            }
        }
        v15 := GenS1{F0: (^(p0))}
        _ = fi11
        _ = v12
        _ = v15
    }
    var v16 float64 = v5
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            switch ("" + (("foo bar" + "foo bar") + "hello")) {
            case "foo bar":
                v19 := []int{fi18, (^(fi18))}
                _ = v19
            case "\n\t":
                var v20 float64 = 0.1
                _ = v20
            case "a":
                fmt.Println(p0)
            default:
                v21 := []int{(fi17 - 1), (fi17 % ((-9223372036854775808) | 1)), fi18, (fi18 ^ fi17)}
                _ = v21
            }
            _ = fi18
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v5
    return 0
}
func genFn1(p0 uint8, p1 int) int16 {
    var v22 float64 = (0.1 - 1e10)
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := []int{p1, fi23, 9223372036854775807}
        v25 := GenS2{F0: 297, F1: fi23}
        for fi26 := 0; fi26 < 9; fi26++ {
            if (v22 > float64(3.14159)) {
                var v27 int = (p1 - 32768)
                fmt.Println(fi23)
                _ = v27
            } else {
                var v28 uint64 = 844
                p1 = -226
                _ = v28
            }
            var v29 int8 = 100
            v30 := GenS3{F0: 42, F1: -175}
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi26
            _ = v29
            _ = v30
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v22
    return 303
}

func main() {
    var v1 float32 = (3.14159 + (1.5 + (0.0 / (1.5 - ((float32(1e10) / 1.5) - 3.14159)))))
    var v2 float64 = 3.14159
    var v3 int16 = 1000
    var v4 uint8 = 127
    v3 += (v3 - 268)
    v1 = v1
    if (4294967295 <= 2) {
        var v31 int16 = 0
        var v32 uint = 18446744073709551615
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int16 = (v3 & v3)
        fmt.Println(v34)
        v35 := []int{}
        _ = fi33
        _ = v34
        _ = v35
    }
    switch "foo bar" {
    case "\n\t":
        for fi36 := 0; fi36 < 1; fi36++ {
            v3 -= 1
            _ = fi36
        }
    case "":
        var v37 uint16 = 255
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        if ((v3 <= (v3 & v3)) || true) {
            var v39 string = "\n\t"
            fmt.Println(v2)
            var v40 uint64 = 18446744073709551615
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            var v42 bool = false
            var v43 int16 = (v3 ^ v3)
            v44 := GenS3{F0: 255, F1: (v3 / ((genFn1(v4, (fi41 ^ fi38))) | 1))}
            var v45 uint = 472
            v1 += 1e10
            v46 := []int{2147483647}
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
            _ = v45
            _ = v46
        }
        _ = fi38
    }
    if ((v4 >> ((65535) & 63)) >= v4) {
        v2 += 0.0
    }
    var v47 int8 = 89
    var v48 int = -966
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
}
