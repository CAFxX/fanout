package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0(p0 uint32) uint16 {
    var v7 uint = 255
    var v8 uint = (v7 * uint(p0))
    v9 := []int{-408}
    var v10 uint = 18446744073709551615
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 32767
}
func genFn1() string {
    for fi11 := 0; fi11 < 19; fi11++ {
        if ((true || (4294967295 != 686)) && (!(true || false))) {
            var v12 float32 = 3.14159
            _ = v12
        } else {
            var v13 float64 = 1e10
            v13 -= 0.0
            _ = v13
        }
        _ = fi11
    }
    v14 := []int{32767, 9223372036854775807}
    if false {
        for fi15 := 0; fi15 < 15; fi15++ {
            if true {
                if (false && (false || false)) {
                    var v16 uint32 = 65535
                    v17 := GenS1{F0: (("" + "foo bar") + "foo bar")}
                    switch "\n\t" {
                    case "hello":
                        if (!(("a" + "hello") == ("\n\t" + "hello"))) {
                            if true {
                                v14 = append(v14, fi15)
                            }
                            var v18 uint8 = 252
                            v19 := GenS1{F0: "hello"}
                            _ = v18
                            _ = v19
                        }
                    case "a":
                        v14 = append(v14, 663)
                    default:
                        fmt.Println(v16)
                    }
                    _ = v16
                    _ = v17
                } else {
                    v20 := GenS1{F0: "a"}
                    _ = v20
                }
            } else {
                var v21 string = ""
                v21 = v21
                _ = v21
            }
            _ = fi15
        }
    } else {
        v22 := GenS1{F0: ("foo bar" + ("\n\t" + "\n\t"))}
        _ = v22
    }
    switch "\n\t" {
    case "hello":
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: "\n\t"}
            v25 := GenS1{F0: "hello"}
            _ = len(v14)
            _ = fi23
            _ = v24
            _ = v25
        }
    case "foo bar":
        for fi26 := 0; fi26 < 4; fi26++ {
            switch ("hello" + ("" + ("\n\t" + "foo bar"))) {
            case "foo bar":
                if false {
                    v27 := GenS2{F0: -1, F1: "foo bar", F2: fi26}
                    if (true || false) {
                        if len(v14) > 0 {
                            v14[0] = fi26
                        }
                    } else {
                        var v28 uint8 = 255
                        _ = v28
                    }
                    fmt.Println(fi26)
                    _ = v27
                }
            }
            _ = fi26
        }
    default:
        if false {
            var v29 uint8 = 255
            var v30 bool = (1e10 != 1.5)
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 int32 = 42
        if true {
            fmt.Println(v32)
        } else {
            v33 := GenS2{F0: (-(v32)), F1: "a", F2: fi31}
            v32 = 464
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    _ = v14
    return "foo bar"
}

func main() {
    var v1 int64 = 624
    var v2 int8 = 127
    var v3 bool = (false || false)
    var v4 bool = v3
    v5 := make(map[uint32]int)
    var v6 int32 = 2147483647
    switch (("hello" + "") + "a") {
    case "a":
        var v34 uint64 = 1
        _ = v34
    case "\n\t":
        var v35 int32 = v6
        _ = v35
    case "foo bar":
        if v4 {
            v3 = v4
        } else {
            delete(v5, 2)
        }
    default:
        switch -9223372036854775808 {
        case -235:
            v4 = (-445 != -129)
        case -2147483648:
            var v36 uint64 = 336
            _ = v36
        case 1:
            var v37 string = (("" + ("hello" + "\n\t")) + "a")
            _ = v37
        }
    }
    if false {
        for fi38 := 0; fi38 < 10; fi38++ {
            fmt.Println(v4, v6)
            var v39 bool = v4
            fmt.Println(v6)
            _ = fi38
            _ = v39
        }
    } else {
        for fi40 := 0; fi40 < 17; fi40++ {
            v2 -= v2
            v6 = v6
            _ = fi40
        }
    }
    fmt.Println(v6, v6)
    if (691 <= 65535) {
        var v41 float64 = (((1.5 * (1e10 + 0.0)) - 0.1) * -2.25)
        _ = v41
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        v1 += v1
        v4 = v3
        _ = fi42
    }
    var v43 uint16 = 990
    v3 = (42 > 2)
    v44 := GenS3{F0: int16(v43)}
    var v45 int32 = (v6 % (((^(v6))) | 1))
    fmt.Println(v43, v4)
    v46 := GenS4{F0: 1, F1: int8(v1), F2: 1}
    var v47 int8 = v2
    var v48 string = ((("\n\t" + "hello") + (("hello" + "a") + ("foo bar" + ""))) + ("" + ("foo bar" + "a")))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
