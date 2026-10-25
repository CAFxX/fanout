package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 uint32, p1 uint16) int8 {
    var v6 string = "hello"
    for fi7 := 0; fi7 < 6; fi7++ {
        switch v6 {
        case "":
            var v8 int = fi7
            _ = v8
        default:
            if false {
                if (!false) {
                    if false {
                        switch fi7 {
                        case 103:
                            var v9 int16 = -1
                            _ = v9
                        case -9223372036854775808:
                            p0 = (p0 & 224)
                        }
                    }
                }
            } else {
                var v10 uint8 = 255
                _ = v10
            }
        }
        v11 := GenS1{F0: 1, F1: 18446744073709551615}
        v12 := GenS1{F0: -9223372036854775808, F1: 141}
        _ = fi7
        _ = v11
        _ = v12
    }
    if (p1 <= (^(p1))) {
        fmt.Println(v6, p1)
    } else {
        if (2 < 4294967295) {
            var v13 int = -2147483648
            if false {
                v14 := GenS1{F0: 42, F1: 42}
                var v15 uint = 42
                v16 := GenS2{F0: 42, F1: "a"}
                _ = v14
                _ = v15
                _ = v16
            }
            _ = v13
        }
        fmt.Println(v6, p1)
        v17 := GenS1{F0: -310, F1: 993}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v6
    return 100
}
func genFn1(p0 string, p1 uint8) uint8 {
    var v18 int64 = 0
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint32 = 2147483647
        var v21 string = ""
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 1, F1: 2}
    var v23 int64 = v18
    _ = p0
    _ = p1
    _ = v18
    _ = v22
    _ = v23
    return (p1 * p1)
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    var v3 uint16 = 32767
    var v4 string = "a"
    var v5 int64 = 0
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS1{F0: (^(v5)), F1: 0}
        _ = fi24
        _ = v25
    }
    switch (v4 + v4) {
    case "hello":
        if false {
            switch 128 {
            case 9223372036854775807:
                v5 = v5
            case -998:
                switch 678 {
                case 0:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v27 := GenS2{F0: -59, F1: ((v4 + "\n\t") + "\n\t")}
                        if false {
                            switch "\n\t" {
                            case "hello":
                                v1[""] = true
                            }
                            v28 := []int{(fi26 | (fi26 - 32768))}
                            _ = v28
                        } else {
                            v2 = (v2 >> (((v2 + 2)) & 63))
                            v4 = (v4 + ("hello" + ("foo bar" + ("a" + v4))))
                        }
                        var v29 int32 = -178
                        _ = fi26
                        _ = v27
                        _ = v29
                    }
                case -843:
                    _ = len(v1)
                }
            default:
                v30 := []int{42, 128}
                _ = v30
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 int64 = v5
                v33 := GenS3{F0: v4, F1: 1000, F2: 0}
                _ = fi31
                _ = v32
                _ = v33
            }
            if true {
                v3 -= 255
                v34 := GenS2{F0: genFn0(2147483648, 493), F1: "a"}
                _ = v34
            }
        }
    case "\n\t":
        v3 = v3
    }
    var v35 bool = (2147483647 == 2147483647)
    for fi36 := 0; fi36 < 10; fi36++ {
        if false {
            for fi37 := 0; fi37 < 7; fi37++ {
                var v38 string = "\n\t"
                _ = fi37
                _ = v38
            }
            var v39 int8 = 87
            _ = v39
        } else {
            v40 := GenS4{F0: 971, F1: 1}
            switch ("hello" + (v4 + v4)) {
            case "":
                switch fi36 {
                case 42:
                    v35 = (false && (!("a" < ("\n\t" + v4))))
                case -1:
                    var v41 uint32 = 342
                    _ = v41
                case 127:
                    if false {
                        var v42 int8 = 42
                        fmt.Println(v42, v5, v3)
                        var v43 uint8 = genFn1(v4, 1)
                        _ = v42
                        _ = v43
                    }
                default:
                    for fi44 := 0; fi44 < 11; fi44++ {
                        var v45 uint32 = 424
                        _ = fi44
                        _ = v45
                    }
                }
            default:
                if false {
                    var v46 uint8 = 2
                    _ = v46
                }
            }
            _ = v40
        }
        v47 := GenS1{F0: (v5 - -1), F1: v2}
        _ = fi36
        _ = v47
    }
    delete(v1, (v4 + "a"))
    var v48 int16 = 484
    v5 += v5
    fmt.Println(v5, v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v48
    _ = v5
}
