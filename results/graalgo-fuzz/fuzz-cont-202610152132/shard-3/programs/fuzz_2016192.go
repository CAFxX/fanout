package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = true
    var v3 int16 = 32767
    var v4 uint32 = 2147483647
    fmt.Println(v4, v4)
    v3 -= (v3 - (v3 + 0))
    v1[4294967295] = "foo bar"
    v3 = v3
    v5 := GenS1{F0: 1}
    v6 := GenS1{F0: 145}
    var v7 bool = true
    v8 := GenS1{F0: 127}
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            switch (fi9 ^ (fi10 - (^(fi10)))) {
            case 32768:
                v11 := GenS1{F0: 2}
                _ = v11
            case 128:
                switch "a" {
                case "foo bar":
                    var v12 uint32 = (v4 / ((82) | 1))
                    var v13 int32 = 0
                    v1[v4] = "hello"
                    _ = v12
                    _ = v13
                case "":
                    _ = len(v1)
                }
            }
            v2 = false
            _ = fi10
        }
        if ((!("foo bar" <= "")) && v7) {
            var v14 uint16 = 0
            var v15 uint16 = v14
            _ = v14
            _ = v15
        } else {
            if (-2.25 >= -2.25) {
                fmt.Println(v3, fi9)
                fmt.Println(v3, v3)
                for fi16 := 0; fi16 < 19; fi16++ {
                    v7 = ((v3 % (((v3 * 1)) | 1)) >= (v3 << ((2) & 63)))
                    _ = fi16
                }
            }
            v4 = v4
            v2 = (((113 != 197) && ((fi9 < 0) && true)) || v7)
        }
        v17 := GenS1{F0: 115}
        _ = fi9
        _ = v17
    }
    if v7 {
        v4 -= 2147483648
        v7 = ((1 <= 18446744073709551615) && (477 >= 65535))
    } else {
        v18 := []int{9223372036854775807, 4294967295, 127}
        _ = v18
    }
    if (42 != 1) {
        v19 := GenS1{F0: 1}
        if (!v7) {
            _ = len(v1)
            for fi20 := 0; fi20 < 4; fi20++ {
                switch (fi20 + fi20) {
                case -9223372036854775808:
                    if false {
                        var v21 uint8 = 127
                        _ = v21
                    }
                case -2147483648:
                    switch (("a" + ("hello" + (("hello" + "hello") + "\n\t"))) + "hello") {
                    case "a":
                        v3 += 1000
                    case "":
                        if (2 > v4) {
                            v22 := []int{fi20, -2147483648, fi20, (fi20 >> ((42) & 63))}
                            var v23 bool = v7
                            _ = v22
                            _ = v23
                        }
                    case "\n\t":
                        v24 := v1[v4]
                        _ = v24
                    default:
                        v25 := []int{(fi20 - fi20), fi20, (fi20 & (fi20 - (fi20 ^ fi20)))}
                        _ = v25
                    }
                default:
                    switch fi20 {
                    case 450:
                        v26 := []int{103, fi20, (fi20 ^ 127), (fi20 % ((4294967295) | 1))}
                        _ = v26
                    case 580:
                        var v27 uint8 = 42
                        _ = v27
                    case -2147483648:
                        var v28 uint64 = 2
                        _ = v28
                    default:
                        if v2 {
                            var v29 uint = 350
                            _ = v29
                        } else {
                            v3 = v3
                            fmt.Println(v7, v3, v3)
                            var v30 float32 = (0.1 / 1e10)
                            _ = v30
                        }
                    }
                }
                if ((fi20 & (fi20 >> ((42) & 63))) >= fi20) {
                    v4 = (v4 & v4)
                    v7 = (73 != 127)
                    v4 += v4
                }
                for fi31 := 0; fi31 < 14; fi31++ {
                    v4 = (^(v4))
                    fmt.Println(fi20)
                    var v32 int8 = 100
                    _ = fi31
                    _ = v32
                }
                _ = fi20
            }
            v33 := []int{-707}
            _ = v33
        }
        var v34 string = ""
        _ = v19
        _ = v34
    }
    if ((381 > 68) || false) {
        v35 := GenS1{F0: 255}
        for fi36 := 0; fi36 < 20; fi36++ {
            v3 += v3
            _ = len(v1)
            _ = fi36
        }
        _ = v35
    }
    fmt.Println(v4, v3)
    v4 = (v4 & v4)
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 bool = true
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
