package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64 }
func genFn0() uint64 {
    var v5 int = 0
    fmt.Println(v5)
    switch v5 {
    case 32767:
        v5 -= (v5 & v5)
    }
    v5 -= (v5 & (^(v5)))
    var v6 uint = 4294967295
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1() bool {
    var v7 uint32 = 187
    for fi8 := 0; fi8 < 19; fi8++ {
        switch ("a" + ("" + "foo bar")) {
        case "hello":
            if ((("foo bar" + "\n\t") + ("\n\t" + ((("\n\t" + "") + "a") + "a"))) != "") {
                var v9 float64 = 3.14159
                for fi10 := 0; fi10 < 2; fi10++ {
                    if (!true) {
                        var v11 float32 = ((0.0 + 1.5) + -2.25)
                        v12 := GenS1{F0: 4294967295, F1: "hello"}
                        _ = v11
                        _ = v12
                    } else {
                        v13 := []int{(fi10 + fi8), fi8}
                        v7 += v7
                        _ = v13
                    }
                    var v14 int64 = 68
                    v15 := []int{}
                    var v16 float32 = (-2.25 - float32((-2.25 + (v9 - (v9 - v9)))))
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi10
                    _ = v14
                    _ = v15
                    _ = v16
                }
                v17 := []int{(fi8 - fi8)}
                _ = v17
                _ = v9
            } else {
                v7 = (v7 ^ v7)
                v7 = v7
            }
        default:
            v18 := GenS1{F0: 4294967295, F1: "\n\t"}
            _ = v18
        }
        var v19 uint32 = (v7 << ((65535) & 63))
        v20 := []int{fi8, 32767, fi8, (fi8 / ((fi8) | 1))}
        _ = fi8
        _ = v19
        _ = v20
    }
    _ = v7
    return (3.14159 < ((0.1 + ((float64(0.1) - (1.5 + 1e10)) * 0.1)) / 0.1))
}

func main() {
    var v1 int64 = 42
    var v2 uint64 = 1
    var v3 string = "hello"
    v4 := []string{}
    var v21 int32 = 2147483647
    switch (((v3 + (v3 + v3)) + "\n\t") + v3) {
    case "":
        v3 = v3
    case "a":
        v21 -= v21
    default:
        var v22 int64 = (v1 | 0)
        _ = v22
    }
    v21 -= (v21 & -681)
    if (((v3 + (v3 + (v3 + v3))) + v3) <= "foo bar") {
        var v23 string = v3
        _ = v23
    } else {
        if len(v4) > 0 {
            v4[0] = v3
        }
    }
    switch 830 {
    case 378:
        if false {
            var v24 int = 32767
            _ = v24
        }
    default:
        var v25 int8 = 0
        _ = v25
    }
    v1 -= (v1 >> ((111) & 63))
    fmt.Println(v1, v21, v21)
    if (32767 != 1) {
        switch -353 {
        case 2147483647:
            fmt.Println(v3)
        case 667:
            if (true && true) {
                var v26 int32 = (-(v21))
                for fi27 := 0; fi27 < 15; fi27++ {
                    v3 = v3
                    for fi28 := 0; fi28 < 10; fi28++ {
                        switch (fi28 / (((^(fi27))) | 1)) {
                        case -9223372036854775808:
                            v4 = append(v4, v3)
                        case 2147483647:
                            var v29 uint16 = 230
                            _ = v29
                        case 9223372036854775807:
                            var v30 bool = true
                            _ = v30
                        }
                        var v31 uint64 = (v2 >> ((2) & 63))
                        _ = fi28
                        _ = v31
                    }
                    _ = fi27
                }
                _ = v26
            }
        case 9223372036854775807:
            v2 = (^(v2))
        default:
            switch 1 {
            case -9223372036854775808:
                v4 = append(v4, v3)
            }
        }
    } else {
        v32 := GenS1{F0: 65535, F1: ("" + "foo bar")}
        _ = v32
    }
    v2 = v2
    switch 941 {
    case 1:
        switch 1 {
        case -306:
            v1 -= v1
        case -860:
            v33 := GenS1{F0: 104, F1: ""}
            _ = v33
        default:
            if (v1 < v1) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v21 = -800
                var v34 float64 = 1.5
                _ = v34
            } else {
                var v35 uint32 = 1
                v36 := GenS1{F0: 4294967295, F1: v3}
                if true {
                    if ((!(42 > 255)) && (((true || false) && false) || (9223372036854775807 == -710))) {
                        if ((v1 & v1) >= (v1 | -1)) {
                            _ = len(v4)
                        }
                    }
                    for fi37 := 0; fi37 < 20; fi37++ {
                        fmt.Println(v2, v3)
                        _ = fi37
                    }
                } else {
                    fmt.Println(v35, v2, v21)
                    v4 = append(v4, "")
                }
                _ = v35
                _ = v36
            }
        }
    case 127:
        switch ("hello" + v3) {
        case "\n\t":
            v38 := GenS1{F0: 42, F1: "hello"}
            _ = v38
        default:
            var v39 int8 = 42
            _ = v39
        }
    }
    var v40 int = -1
    v41 := GenS2{F0: v1}
    v42 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
