package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint32) uint8 {
    if false {
        v7 := GenS1{F0: -32768, F1: 0, F2: 644}
        _ = v7
    } else {
        switch "hello" {
        case "hello":
            if (!(!true)) {
                p1 = p1
                fmt.Println(p1)
            } else {
                if false {
                    if ((425 < 436) && true) {
                        var v8 uint16 = 255
                        _ = v8
                    }
                }
                v9 := GenS1{F0: 42, F1: 1, F2: 65535}
                _ = v9
            }
        case "":
            fmt.Println(p0, p0)
        case "a":
            switch "\n\t" {
            case "\n\t":
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS2{F0: 231}
                    _ = fi10
                    _ = v11
                }
            default:
                p0 = 42
            }
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 float64 = ((-2.25 - -2.25) * 1e10)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    if (true || true) {
        p1 += p1
    } else {
        var v15 uint = 4294967295
        for fi16 := 0; fi16 < 18; fi16++ {
            switch (("" + ("foo bar" + ("a" + "foo bar"))) + "\n\t") {
            case "\n\t":
                var v17 int64 = (p0 + (p0 << ((2) & 63)))
                _ = v17
            }
            switch (fi16 * (fi16 + (-(fi16)))) {
            case 2147483647:
                v18 := GenS1{F0: -32768, F1: (-(fi16)), F2: (v15 - (v15 | v15))}
                _ = v18
            default:
                if (true && (1.5 != 0.1)) {
                    var v19 int16 = 874
                    _ = v19
                }
            }
            var v20 uint64 = 685
            _ = fi16
            _ = v20
        }
        _ = v15
    }
    switch ("hello" + "") {
    case "\n\t":
        var v21 int32 = 2147483647
        _ = v21
    case "":
        var v22 uint64 = 4294967295
        _ = v22
    case "a":
        var v23 uint16 = 32767
        _ = v23
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint16, p1 uint8, p2 uint32) uint64 {
    p2 -= p2
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint64 = 4294967295
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        p2 = 1
        _ = fi26
    }
    var v27 uint16 = uint16(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 604
}

func main() {
    var v1 int32 = -1
    var v2 string = "a"
    var v3 bool = (255 <= 32767)
    v4 := []string{}
    v5 := make(map[int]int)
    var v6 uint8 = 1
    var v28 int8 = -60
    var v29 int32 = (v1 | (v1 << ((42) & 63)))
    v30 := GenS1{F0: 32767, F1: 0, F2: 0}
    v6 = (v6 % (((v6 | v6)) | 1))
    var v31 uint64 = 42
    v1 += 1
    if true {
        fmt.Println(v6, v3, v2)
    } else {
        var v32 uint32 = 4294967295
        v28 = v28
        v33 := GenS1{F0: 42, F1: 1, F2: uint(v31)}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        switch fi34 {
        case 127:
            v4 = append(v4, v2)
        default:
            switch int(v28) {
            case -2147483648:
                v3 = (168 == 822)
            default:
                v2 = v2
            }
        }
        var v35 uint64 = 2
        _ = fi34
        _ = v35
    }
    if v3 {
        v31 = v31
        for fi36 := 0; fi36 < 11; fi36++ {
            v4 = append(v4, (v2 + "hello"))
            if (!(v3 || v3)) {
                if true {
                    var v37 int16 = -32768
                    var v38 int64 = int64(v37)
                    _ = v37
                    _ = v38
                }
                fmt.Println(v31)
            } else {
                var v39 float32 = (((float32(float64(3.14159)) - ((1.5 * 0.0) + -2.25)) + 1.5) / -2.25)
                v31 = v31
                for fi40 := 0; fi40 < 9; fi40++ {
                    _ = len(v5)
                    var v41 int16 = 245
                    switch (fi40 * (fi40 | fi36)) {
                    case -1:
                        var v42 int64 = 42
                        _ = v42
                    case 127:
                        v1 -= v1
                    case 879:
                        if v3 {
                            var v43 int32 = -951
                            _ = v43
                        } else {
                            v6 += (v6 >> ((18446744073709551615) & 63))
                            v39 -= 3.14159
                        }
                    }
                    delete(v5, (fi40 * fi40))
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
            _ = fi36
        }
    }
    switch 127 {
    case -2147483648:
        if ((!(v3 || v3)) && (v1 != v1)) {
            var v44 uint8 = 128
            _ = v44
        } else {
            v45 := GenS2{F0: 96}
            var v46 bool = v3
            var v47 int16 = 42
            _ = v45
            _ = v46
            _ = v47
        }
    case 42:
        fmt.Println(v3, v31, v6)
    default:
        _ = len(v4)
    }
    v48 := GenS2{F0: 2}
    for fi49 := 0; fi49 < 10; fi49++ {
        v3 = false
        _ = fi49
    }
    v3 = ((3.14159 - float64(1e10)) <= 0.1)
    switch 1 {
    case -1:
        var v50 uint8 = (v6 / ((genFn0(715, 372)) | 1))
        _ = v50
    }
    var v51 float64 = 3.14159
    var v52 int8 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
