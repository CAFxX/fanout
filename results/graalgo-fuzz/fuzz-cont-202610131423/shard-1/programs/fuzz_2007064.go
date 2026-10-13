package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 string }
func genFn0(p0 int8, p1 bool, p2 uint16) int64 {
    p2 -= (p2 << ((104) & 63))
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "foo bar", F1: 1000}
    if (2 == 127) {
        v10 := GenS2{F0: 42, F1: 2, F2: -900}
        _ = v10
    }
    if p1 {
        var v11 int32 = 999
        _ = v11
    } else {
        v12 := GenS1{F0: (("a" + ("a" + "foo bar")) + "hello"), F1: 32767}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 9223372036854775807
}
func genFn1() uint16 {
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 string = "a"
        v14 = "hello"
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int64 = -9223372036854775808
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch fi17 {
        case 230:
            var v18 uint16 = 2
            _ = v18
        default:
            switch fi17 {
            case 317:
                if false {
                    var v19 bool = (995 == 460)
                    var v20 uint64 = 145
                    v19 = (v20 == 579)
                    _ = v19
                    _ = v20
                }
            }
        }
        v21 := GenS2{F0: 802, F1: 694, F2: 9223372036854775807}
        _ = fi17
        _ = v21
    }
    return 42
}

func main() {
    var v1 bool = (((true && true) || (18446744073709551615 <= 42)) || true)
    v2 := []int{}
    v3 := []string{"a", (("foo bar" + "a") + "hello"), ("" + (("hello" + (("foo bar" + "\n\t") + "")) + "\n\t"))}
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{681}
    var v6 int64 = -928
    v2 = append(v2, 32767)
    if true {
        var v22 int8 = 0
        switch ("foo bar" + "") {
        case "foo bar":
            v23 := GenS3{F0: 0, F1: -1}
            _ = v23
        default:
            _ = len(v2)
        }
        v22 = (v22 & v22)
        _ = v22
    } else {
        var v24 float64 = -2.25
        _ = v24
    }
    if true {
        var v25 int64 = (v6 - v6)
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 uint64 = 18446744073709551615
        v6 += 42
        v28 := GenS3{F0: -824, F1: -692}
        _ = fi26
        _ = v27
        _ = v28
    }
    switch "" {
    case "a":
        fmt.Println(v6, v1, v1)
    case "foo bar":
        fmt.Println(v1, v4, v1)
    }
    v6 -= (v6 >> ((139) & 63))
    if len(v2) > 0 {
        v2[0] = -417
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6, v1, v1)
    v3 = append(v3, "")
    _ = len(v2)
    fmt.Println(v4)
    switch -1 {
    case 42:
        if v1 {
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 int64 = 839
                for fi31 := 0; fi31 < 17; fi31++ {
                    v32 := GenS4{F0: 0, F1: "hello"}
                    v1 = true
                    _ = fi31
                    _ = v32
                }
                var v33 int32 = 42
                _ = fi29
                _ = v30
                _ = v33
            }
            var v34 uint32 = 409
            v1 = (true && v1)
            _ = v34
        } else {
            if (!(2 != 127)) {
                _ = len(v3)
                v3 = append(v3, ("foo bar" + ("hello" + ((("foo bar" + "\n\t") + "hello") + "\n\t"))))
                if len(v5) > 0 {
                    v5[0] = 1
                }
            } else {
                v35 := GenS5{F0: uint32(v4), F1: ""}
                switch "foo bar" {
                case "hello":
                    var v36 int64 = (v6 & v6)
                    _ = v36
                case "foo bar":
                    var v37 float64 = 0.0
                    _ = v37
                }
                _ = v35
            }
            if v1 {
                v38 := GenS6{F0: "hello"}
                for fi39 := 0; fi39 < 17; fi39++ {
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        if true {
                            fmt.Println(fi39, fi39, v4)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = (0 >= 2)
                        }
                    }
                    fmt.Println(v6, v1, v6)
                    _ = fi39
                }
                _ = v38
            } else {
                v6 += v6
                v1 = false
                switch 888 {
                case -1:
                    var v40 int32 = 835
                    _ = v40
                case -641:
                    v4 = v4
                default:
                    var v41 int = 2147483647
                    _ = v41
                }
            }
            var v42 float32 = ((1.5 - 0.0) - 3.14159)
            _ = v42
        }
    case 32768:
        var v43 float64 = (0.1 * 1e10)
        _ = v43
    case 891:
        v6 -= (v6 ^ 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
