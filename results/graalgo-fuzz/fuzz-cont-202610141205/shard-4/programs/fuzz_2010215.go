package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 string) uint32 {
    var v5 int32 = p0
    if false {
        var v6 int = -1
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 249
}
func genFn1(p0 int16, p1 int8, p2 uint64) int32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: false}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        p0 -= p0
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 bool = false
    var v2 int = 9223372036854775807
    var v3 uint64 = 598
    var v4 int64 = 0
    switch "hello" {
    case "":
        var v10 float64 = (3.14159 - (-2.25 / 0.1))
        _ = v10
    case "a":
        fmt.Println(v4, v2, v2)
    case "hello":
        var v11 string = ""
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v4 = v4
        for fi13 := 0; fi13 < 6; fi13++ {
            if ("a" >= "a") {
                v4 -= (v4 + (v4 << ((2) & 63)))
                if (!(v1 || v1)) {
                    switch "\n\t" {
                    case "a":
                        v4 = (v4 / (((v4 - (v4 / ((9223372036854775807) | 1)))) | 1))
                        if true {
                            v4 = -1
                            var v14 int8 = -128
                            _ = v14
                        }
                        var v15 float32 = -2.25
                        _ = v15
                    }
                    switch -2147483648 {
                    case 4294967295:
                        v2 = 32768
                    }
                } else {
                    if v1 {
                        var v16 int32 = 1
                        switch fi12 {
                        case 32767:
                            var v17 int32 = v16
                            _ = v17
                        case 1:
                            var v18 uint32 = 42
                            _ = v18
                        case -2147483648:
                            v2 += 127
                        }
                        _ = v16
                    }
                    if v1 {
                        if (uint8(v3) == 128) {
                            var v19 int32 = 2147483647
                            _ = v19
                        } else {
                            var v20 float32 = 0.0
                            _ = v20
                        }
                    }
                }
            } else {
                v21 := []int{v2, fi13, fi13, fi12}
                _ = v21
            }
            _ = fi13
        }
        _ = fi12
    }
    v2 -= v2
    if (172 == 42) {
        if ((v3 & v3) != (v3 + 989)) {
            var v22 uint32 = 374
            var v23 bool = (("foo bar" + "") == "hello")
            switch (v2 + (v2 * v2)) {
            case 200:
                fmt.Println(v1)
            }
            _ = v22
            _ = v23
        }
    }
    var v24 uint8 = 127
    var v25 float32 = float32(0.1)
    v4 = v4
    var v26 int32 = genFn1(597, 1, 4294967295)
    switch "a" {
    case "":
        v24 = 144
    case "hello":
        if (true && ((!(!((v1 && (!("a" >= "foo bar"))) || true))) && (!(!v1)))) {
            v4 -= (v4 + v4)
            switch (^(v2)) {
            case -457:
                var v27 int64 = v4
                _ = v27
            case 4294967295:
                v28 := GenS1{F0: (("" <= "a") && v1)}
                _ = v28
            default:
                var v29 uint16 = 2
                _ = v29
            }
        }
    case "foo bar":
        var v30 bool = v1
        _ = v30
    }
    v31 := GenS1{F0: (!true)}
    if true {
        switch (v2 & v2) {
        case 4294967295:
            v32 := GenS1{F0: (!false)}
            _ = v32
        case -9223372036854775808:
            var v33 uint16 = 413
            _ = v33
        default:
            v24 += v24
        }
        v34 := []int{}
        _ = v34
    }
    v35 := GenS2{F0: 394}
    if false {
        for fi36 := 0; fi36 < 1; fi36++ {
            v37 := []int{fi36, (-(fi36))}
            var v38 uint64 = (v3 & v3)
            var v39 uint = 2
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
    }
    var v40 uint64 = v3
    v2 = v2
    var v41 float32 = (-2.25 + v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v40
    _ = v41
}
