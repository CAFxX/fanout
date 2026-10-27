package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
func genFn0() int16 {
    v6 := []int{128, 2147483647, 312}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 bool = (false || (!(fi7 == fi7)))
        switch (fi7 | (fi7 * (-(fi7)))) {
        case -314:
            var v9 int64 = 1
            _ = v9
        case -204:
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int8 = 0
                v11 += (-(v11))
                _ = fi10
                _ = v11
            }
        default:
            v12 := GenS1{F0: 9223372036854775807}
            _ = v12
        }
        switch "hello" {
        case "hello":
            v8 = ((((65535 <= 2) || (!true)) && true) || ("\n\t" == "\n\t"))
        case "a":
            v8 = (-414 == fi7)
        case "\n\t":
            if v8 {
                _ = len(v6)
                v8 = (!true)
            } else {
                v8 = v8
            }
        }
        _ = fi7
        _ = v8
    }
    var v13 uint = 1
    _ = v13
    _ = v6
    return 1000
}
func genFn1() uint {
    var v14 int = 511
    v14 += (v14 >> ((0) & 63))
    v15 := GenS1{F0: 9223372036854775807}
    var v16 int64 = -479
    _ = v14
    _ = v15
    _ = v16
    return 18446744073709551615
}

func main() {
    var v1 int = 107
    var v2 uint8 = 255
    var v3 uint8 = v2
    var v4 uint = 163
    v5 := []string{("" + ""), ("foo bar" + "foo bar"), "hello", ("foo bar" + "foo bar")}
    var v17 string = ("a" + ((("\n\t" + ("hello" + (("\n\t" + "") + "a"))) + (("hello" + "") + ("" + "a"))) + (("\n\t" + (("\n\t" + "hello") + "hello")) + "hello")))
    if (255 == 597) {
        switch "hello" {
        case "a":
            fmt.Println(v1, v17, v17)
        default:
            var v18 int8 = -104
            _ = v18
        }
        var v19 string = "hello"
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint16 = 949
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := GenS1{F0: 201}
                if (true || true) {
                    var v24 uint8 = (v3 + 2)
                    switch v17 {
                    case "hello":
                        v19 = ""
                    }
                    switch (v17 + v17) {
                    case "foo bar":
                        var v25 uint32 = 2147483648
                        _ = v25
                    default:
                        var v26 int32 = 2147483647
                        _ = v26
                    }
                    _ = v24
                } else {
                    v17 = v17
                    var v27 uint8 = v2
                    _ = v27
                }
                v2 += v2
                if true {
                    var v28 float64 = 1.5
                    v29 := GenS1{F0: 190}
                    _ = v28
                    _ = v29
                } else {
                    v19 = v17
                    v30 := GenS1{F0: -757}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v30
                }
                _ = fi22
                _ = v23
            }
            var v31 int8 = 1
            _ = fi20
            _ = v21
            _ = v31
        }
        _ = v19
    } else {
        var v32 uint32 = 2
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        if true {
            var v34 int = (fi33 | (v1 / (((fi33 - (fi33 % (((-(fi33))) | 1)))) | 1)))
            switch v17 {
            case "hello":
                v35 := GenS2{F0: -1}
                _ = v35
            case "foo bar":
                fmt.Println(v1, v17)
            }
            var v36 string = v17
            _ = v34
            _ = v36
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 uint32 = 65535
                _ = fi37
                _ = v38
            }
            var v39 int = fi33
            _ = v39
        }
        switch ("a" + v17) {
        case "hello":
            var v40 bool = true
            _ = v40
        case "\n\t":
            v2 = v3
        case "a":
            if true {
                v1 = 839
                var v41 string = "\n\t"
                _ = v41
            }
        default:
            var v42 uint16 = 414
            _ = v42
        }
        v4 = genFn1()
        _ = fi33
    }
    v43 := GenS2{F0: 1}
    var v44 int16 = -549
    var v45 int = (v1 ^ -52)
    v46 := GenS3{F0: 32768}
    v47 := GenS4{F0: (v2 % ((uint8(v4)) | 1))}
    v44 = -1
    if (!true) {
        if len(v5) > 0 {
            v5[0] = v17
        }
        var v48 float32 = 0.0
        _ = v48
    } else {
        v44 -= -32768
        var v49 uint = (v4 % ((v4) | 1))
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
