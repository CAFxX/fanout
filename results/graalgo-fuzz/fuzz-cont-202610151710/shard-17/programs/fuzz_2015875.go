package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int32 {
    if (0.0 != 0.0) {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: 1}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 bool = (!(18446744073709551615 < 18446744073709551615))
                v12 := []int{(^(fi10)), 1}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        if (false && true) {
            v13 := []int{9223372036854775807}
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v15 string = "foo bar"
                _ = v15
            case "foo bar":
                if true {
                    var v16 int8 = -128
                    var v17 float32 = 1e10
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        switch (-(fi18)) {
                        case -603:
                            fmt.Println(fi18, fi14)
                        case 664:
                            fmt.Println(fi18, fi14)
                        default:
                            var v19 int64 = 1
                            _ = v19
                        }
                        _ = fi18
                    }
                }
            default:
                fmt.Println(fi14)
            }
            if true {
                if false {
                    var v20 int64 = 0
                    var v21 uint16 = 383
                    _ = v20
                    _ = v21
                }
                var v22 float64 = (-2.25 + 3.14159)
                var v23 uint64 = 18446744073709551615
                _ = v22
                _ = v23
            } else {
                var v24 uint64 = 424
                var v25 int64 = -9223372036854775808
                _ = v24
                _ = v25
            }
            _ = fi14
        }
        var v26 float64 = (1.5 + (1.5 + 1.5))
        _ = v26
    }
    var v27 int64 = 1
    var v28 uint16 = 42
    var v29 uint32 = 2147483648
    _ = v27
    _ = v28
    _ = v29
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 4294967295
    v3 := []int{128, -9223372036854775808, -355, -9223372036854775808}
    var v4 float32 = (0.1 - 1.5)
    v5 := []uint32{287}
    var v6 uint16 = 425
    v4 -= v4
    v6 += (v6 * (^(v6)))
    _ = len(v5)
    if (128 != 232) {
        switch 32767 {
        case -9223372036854775808:
            var v30 uint32 = 0
            _ = v30
        case 2147483647:
            if true {
                var v31 int16 = 941
                switch -2147483648 {
                case 2147483647:
                    v31 -= v31
                case 128:
                    v32 := GenS1{F0: uint8(v2)}
                    _ = v32
                }
                _ = v31
            }
        case -226:
            var v33 float32 = (v4 / 1.5)
            _ = v33
        }
        v6 = 65535
    }
    switch 2147483647 {
    case -205:
        switch ("foo bar" + ("foo bar" + (("foo bar" + "\n\t") + ("foo bar" + "a")))) {
        case "hello":
            v34 := v1["hello"]
            _ = v34
        case "a":
            if ("foo bar" != "") {
                for fi35 := 0; fi35 < 9; fi35++ {
                    fmt.Println(v6, v4)
                    _ = fi35
                }
                var v36 int = -59
                _ = v36
            } else {
                for fi37 := 0; fi37 < 5; fi37++ {
                    fmt.Println(v2, v2)
                    v2 -= v2
                    if (2 > uint8(v2)) {
                        switch "foo bar" {
                        case "":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "hello":
                            var v38 uint16 = v6
                            _ = v38
                        default:
                            fmt.Println(fi37, v6, v6)
                        }
                    } else {
                        var v39 string = "hello"
                        fmt.Println(v4)
                        v40 := v1[v39]
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                }
            }
        case "":
            var v41 uint8 = 0
            _ = v41
        default:
            var v42 int64 = 0
            _ = v42
        }
    case 128:
        var v43 int = int(v2)
        _ = v43
    case -2147483648:
        var v44 int8 = 127
        _ = v44
    }
    var v45 float32 = (float32(1.5) + -2.25)
    var v46 uint = v2
    if (!(!true)) {
        v47 := GenS1{F0: 225}
        fmt.Println(v45)
        _ = v47
    }
    var v48 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
}
