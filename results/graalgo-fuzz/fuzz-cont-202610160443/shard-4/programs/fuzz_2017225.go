package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 int32) int16 {
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            p1 -= (p1 & (p1 * (p1 + (p1 ^ (p1 << ((18446744073709551615) & 63))))))
            if true {
                var v6 int8 = -1
                if false {
                    v7 := []int{-212, (fi5 - 4294967295), fi5}
                    var v8 bool = true
                    var v9 uint16 = 31
                    _ = v7
                    _ = v8
                    _ = v9
                }
                fmt.Println(fi4, p0, p1)
                _ = v6
            }
            p1 += (p1 | p1)
            v10 := GenS1{F0: "\n\t", F1: 996}
            _ = fi5
            _ = v10
        }
        if ((^(fi4)) != (fi4 | (fi4 / (((fi4 - 127)) | 1)))) {
            v11 := []int{623, (fi4 << (((p0 * p0)) & 63)), fi4}
            v11 = append(v11, fi4)
            p0 -= p0
            _ = v11
        }
        _ = fi4
    }
    var v12 uint = (p0 + (p0 * p0))
    p0 = 0
    _ = p0
    _ = p1
    _ = v12
    return 0
}
func genFn1(p0 uint8, p1 uint) int {
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (("" + "a") + ((("" + "hello") + "foo bar") + "hello")) {
        case "hello":
            fmt.Println(p1, p1, fi13)
        case "\n\t":
            var v14 float32 = 3.14159
            _ = v14
        default:
            for fi15 := 0; fi15 < 18; fi15++ {
                p1 += 266
                switch ((((("\n\t" + "foo bar") + ("\n\t" + "foo bar")) + "foo bar") + ("" + "hello")) + "hello") {
                case "":
                    p0 -= 24
                case "a":
                    var v16 int = -505
                    _ = v16
                default:
                    var v17 int = 9223372036854775807
                    _ = v17
                }
                _ = fi15
            }
        }
        _ = fi13
    }
    p1 -= (p1 + (p1 | p1))
    var v18 int = 0
    _ = p0
    _ = p1
    _ = v18
    return v18
}

func main() {
    var v1 uint16 = 1
    var v2 int16 = -32768
    var v3 uint32 = 1
    v19 := GenS2{F0: -9223372036854775808}
    v1 -= (^(v1))
    var v20 int8 = -117
    switch 0 {
    case 32768:
        if true {
            v21 := []int{-213, genFn1(255, 2), 4294967295, 269}
            _ = v21
        } else {
            var v22 uint32 = v3
            v3 -= (v22 >> ((42) & 63))
            _ = v22
        }
    case 1:
        var v23 int32 = -1
        _ = v23
    default:
        fmt.Println(v1, v2)
    }
    var v24 bool = true
    if v24 {
        v20 = (v20 & v20)
        switch genFn1(127, 0) {
        case 32768:
            v24 = (("\n\t" <= "a") && true)
        default:
            var v25 int = genFn1(2, 132)
            _ = v25
        }
    } else {
        v20 += v20
        for fi26 := 0; fi26 < 15; fi26++ {
            v2 = v2
            _ = fi26
        }
        v27 := GenS1{F0: "a", F1: (v1 ^ (^(v1)))}
        _ = v27
    }
    var v28 uint16 = (v1 & 987)
    v29 := []int{2147483647}
    var v30 int8 = (v20 * (v20 ^ -106))
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v31 := GenS3{F0: (v2 ^ (v2 + 0)), F1: 1, F2: (v28 | v28)}
    v2 = (v2 | 42)
    if true {
        var v32 int64 = -9223372036854775808
        var v33 uint32 = 0
        _ = v32
        _ = v33
    } else {
        v29 = append(v29, 105)
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 int32 = 2147483647
            _ = fi34
            _ = v35
        }
        if v24 {
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v24)
            }
            fmt.Println(v20, v2)
        }
    }
    if v24 {
        var v36 float32 = (((3.14159 + float32(3.14159)) / (float32(float64((3.14159 * 0.1))) - ((float32((1e10 + 1e10)) / (1.5 / 1.0)) - (1e10 - (float32(-2.25) - 1.5))))) / 0.1)
        v20 += (v20 + (-(v20)))
        _ = v36
    }
    v29 = append(v29, -354)
    if v24 {
        v37 := GenS4{F0: false}
        _ = v37
    }
    switch -2147483648 {
    case 1:
        if (!v24) {
            for fi38 := 0; fi38 < 18; fi38++ {
                for fi39 := 0; fi39 < 15; fi39++ {
                    var v40 uint8 = 255
                    _ = fi39
                    _ = v40
                }
                if ((1e10 * -2.25) < 3.14159) {
                    v41 := GenS4{F0: ((genFn1(128, 42) != (fi38 ^ (-(fi38)))) || v24)}
                    _ = v41
                }
                _ = fi38
            }
        }
    case 127:
        v42 := GenS2{F0: 42}
        _ = v42
    }
    var v43 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v43
}
