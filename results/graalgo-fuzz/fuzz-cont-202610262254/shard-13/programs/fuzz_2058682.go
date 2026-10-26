package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0() uint32 {
    v7 := []int{-650}
    var v8 float64 = 0.0
    var v9 uint = 1
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}
func genFn1(p0 uint32, p1 int64) int8 {
    p1 = p1
    p1 = (p1 / ((0) | 1))
    var v10 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    var v1 uint64 = 66
    var v2 float32 = 3.14159
    var v3 int = 9223372036854775807
    var v4 int16 = 1000
    var v5 uint = 0
    var v6 int8 = -128
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int64 = 0
        v13 := GenS1{F0: (!true), F1: 77}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = 1e10
    fmt.Println(v4, v5)
    fmt.Println(v4, v2, v4)
    if ((v5 >> ((v5) & 63)) == (^(v5))) {
        v6 -= v6
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            v2 -= v2
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(v1, v2, v4)
                var v17 uint16 = 65535
                _ = fi16
                _ = v17
            }
            if (false || true) {
                v18 := GenS1{F0: (2147483647 != 320), F1: 2}
                _ = v18
            }
            _ = fi15
        }
        var v19 uint8 = 9
        _ = v19
    }
    v5 = v5
    var v20 float32 = float32(0.0)
    v21 := GenS2{F0: 128}
    fmt.Println(v2)
    if (v2 != (1.5 / 1.0)) {
        for fi22 := 0; fi22 < 6; fi22++ {
            if false {
                if false {
                    var v23 uint64 = v1
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v25 := GenS1{F0: (v3 < -1), F1: 12}
                        switch "" {
                        case "a":
                            v26 := GenS1{F0: false, F1: 42}
                            fmt.Println(v20, fi22)
                            var v27 string = ((("hello" + (("" + "\n\t") + "\n\t")) + ("\n\t" + "")) + ("foo bar" + "\n\t"))
                            _ = v26
                            _ = v27
                        default:
                            v28 := GenS1{F0: (!(174 == 53)), F1: 128}
                            _ = v28
                        }
                        var v29 string = ("a" + "hello")
                        _ = fi24
                        _ = v25
                        _ = v29
                    }
                    _ = v23
                }
            } else {
                v4 += 789
                var v30 uint16 = 1
                switch "a" {
                case "\n\t":
                    v6 = v6
                case "foo bar":
                    v31 := []int{fi22, (fi22 + (v3 - 2147483647))}
                    _ = v31
                case "hello":
                    switch (v3 | 2147483647) {
                    case 0:
                        for fi32 := 0; fi32 < 19; fi32++ {
                            v33 := GenS1{F0: true, F1: 127}
                            _ = fi32
                            _ = v33
                        }
                    case -2147483648:
                        for fi34 := 0; fi34 < 18; fi34++ {
                            var v35 int = (fi22 << (((v5 | v5)) & 63))
                            v5 = (v5 + 65535)
                            _ = fi34
                            _ = v35
                        }
                    case 9223372036854775807:
                        v36 := GenS1{F0: false, F1: 208}
                        _ = v36
                    }
                default:
                    if (1.5 == 0.1) {
                        var v37 uint16 = v30
                        v5 = (v5 / ((v5) | 1))
                        v38 := []int{fi22, (-(v3))}
                        _ = v37
                        _ = v38
                    }
                }
                _ = v30
            }
            var v39 uint32 = 199
            var v40 bool = ((^(v3)) > int(v39))
            _ = fi22
            _ = v39
            _ = v40
        }
    }
    v41 := GenS1{F0: true, F1: 218}
    switch "foo bar" {
    case "foo bar":
        v3 = -1
    case "hello":
        v42 := []int{v3, -25}
        _ = v42
    case "":
        v43 := []int{4294967295, (^(v3)), (v3 * v3)}
        _ = v43
    default:
        v44 := []int{}
        _ = v44
    }
    v45 := GenS2{F0: 0}
    var v46 float32 = v2
    if false {
        v47 := GenS1{F0: false, F1: 15}
        v48 := []int{v3, 42, (v3 | (v3 + int(v5)))}
        _ = v47
        _ = v48
    } else {
        var v49 uint32 = 0
        _ = v49
    }
    fmt.Println(v46, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
