package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float32 = ((1.5 - (-2.25 - 1e10)) + 0.1)
    var v2 float32 = 1e10
    var v3 string = ("foo bar" + ("foo bar" + "\n\t"))
    var v4 string = v3
    var v5 uint = 2
    if (false && false) {
        switch 9223372036854775807 {
        case 0:
            for fi6 := 0; fi6 < 14; fi6++ {
                v1 -= (v2 / 1e10)
                if ((127 == -128) && false) {
                    fmt.Println(fi6, v3, v3)
                    var v7 uint32 = 4294967295
                    _ = v7
                } else {
                    var v8 uint32 = 4294967295
                    var v9 bool = (2 >= 65535)
                    _ = v8
                    _ = v9
                }
                _ = fi6
            }
        case 1:
            var v10 float64 = (1e10 / 1e10)
            _ = v10
        case 32767:
            var v11 uint32 = 2147483648
            _ = v11
        }
        v12 := []int{32767, 2147483647}
        _ = v12
    } else {
        fmt.Println(v1)
        fmt.Println(v3, v4)
    }
    var v13 string = "\n\t"
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v3)
        if true {
            var v15 float32 = 1e10
            _ = v15
        } else {
            if true {
                var v16 int32 = 1
                _ = v16
            }
        }
        _ = fi14
    }
    var v17 string = (v13 + v4)
    fmt.Println(v4, v5, v4)
    var v18 uint8 = 127
    fmt.Println(v5, v18, v1)
    fmt.Println(v4, v18, v18)
    switch 1 {
    case -624:
        v5 = (v5 >> (((v5 % (((v5 / (((v5 - 2)) | 1))) | 1))) & 63))
    case 1:
        var v19 uint64 = 4294967295
        _ = v19
    default:
        var v20 uint8 = (^(v18))
        _ = v20
    }
    if (((895 != 920) || (!(false && true))) && true) {
        v17 = ((v13 + "hello") + "a")
    } else {
        fmt.Println(v13)
        v13 = (v13 + ("foo bar" + v3))
        for fi21 := 0; fi21 < 2; fi21++ {
            if (552 < 505) {
                if false {
                    if true {
                        v22 := []int{fi21, (fi21 | fi21), 32767}
                        _ = len(v22)
                        _ = v22
                    } else {
                        v5 -= (v5 & (v5 - 42))
                    }
                    var v23 uint8 = v18
                    _ = v23
                } else {
                    if ((^(v18)) <= v18) {
                        var v24 uint16 = 42
                        v5 = (^(v5))
                        if false {
                            v25 := GenS1{F0: 9223372036854775807}
                            switch (v3 + v13) {
                            case "hello":
                                v5 = v5
                            case "a":
                                v13 = v4
                            }
                            if ((fi21 + -9223372036854775808) == fi21) {
                                v26 := []int{(fi21 ^ 32767)}
                                _ = v26
                            } else {
                                v27 := GenS2{F0: 32767}
                                _ = v27
                            }
                            _ = v25
                        }
                        _ = v24
                    } else {
                        if (false && true) {
                            v28 := GenS1{F0: 9223372036854775807}
                            var v29 uint8 = (v18 & (v18 % ((1) | 1)))
                            var v30 float32 = (1e10 - (1.5 - float32(((0.0 - (1.5 + 1e10)) + float64(1e10)))))
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            v5 += uint(v18)
                            var v31 int32 = 1
                            _ = v31
                        }
                    }
                    v4 = ""
                    v18 = (v18 % ((127) | 1))
                }
                fmt.Println(fi21)
            }
            var v32 int32 = 298
            v33 := GenS2{F0: -1}
            _ = fi21
            _ = v32
            _ = v33
        }
    }
    v34 := GenS1{F0: -688}
    var v35 uint32 = 273
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 uint32 = 34
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
