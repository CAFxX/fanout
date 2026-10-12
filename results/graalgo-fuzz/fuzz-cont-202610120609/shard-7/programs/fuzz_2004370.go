package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }

func main() {
    var v1 uint64 = 395
    var v2 uint64 = 42
    var v3 float64 = ((-2.25 + 0.1) * 1.5)
    var v4 uint16 = 3
    switch 4294967295 {
    case -2147483648:
        v3 += v3
    case 4294967295:
        switch 53 {
        case 923:
            var v5 int32 = 0
            _ = v5
        case 128:
            if (false && ((v2 < 18446744073709551615) || false)) {
                for fi6 := 0; fi6 < 14; fi6++ {
                    v7 := []int{}
                    v7 = append(v7, 127)
                    _ = fi6
                    _ = v7
                }
            }
        case 2147483647:
            var v8 uint = 18446744073709551615
            _ = v8
        default:
            var v9 int8 = 127
            _ = v9
        }
    case 2147483647:
        switch 127 {
        case 4294967295:
            for fi10 := 0; fi10 < 7; fi10++ {
                switch ("hello" + "hello") {
                case "foo bar":
                    if (1000 < -445) {
                        fmt.Println(v2, v2, v2)
                        var v11 string = "a"
                        fmt.Println(v11)
                        _ = v11
                    } else {
                        var v12 int64 = -98
                        var v13 string = "\n\t"
                        var v14 uint32 = 1
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case "":
                    if (((-32768 <= 1000) || true) && false) {
                        if ("a" < "") {
                            var v15 bool = (true || false)
                            switch (^(fi10)) {
                            case -13:
                                v16 := []int{fi10}
                                _ = v16
                            case 756:
                                var v17 float32 = -2.25
                                _ = v17
                            case 704:
                                v18 := []int{}
                                _ = v18
                            default:
                                for fi19 := 0; fi19 < 15; fi19++ {
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v1, v15, v1)
                                    case "foo bar":
                                        if v15 {
                                            var v20 float64 = v3
                                            fmt.Println(fi19)
                                            var v21 uint64 = v2
                                            _ = v20
                                            _ = v21
                                        }
                                    case "":
                                        var v22 int = (fi19 ^ int(v4))
                                        _ = v22
                                    }
                                    fmt.Println(v2, v2, fi19)
                                    _ = fi19
                                }
                            }
                            var v23 int = 127
                            _ = v15
                            _ = v23
                        } else {
                            v2 = v1
                            for fi24 := 0; fi24 < 5; fi24++ {
                                v25 := []int{fi24, 4294967295, fi24}
                                _ = len(v25)
                                _ = fi24
                                _ = v25
                            }
                            v2 = v1
                        }
                    }
                case "a":
                    var v26 bool = false
                    _ = v26
                }
                _ = fi10
            }
        case -9223372036854775808:
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 uint32 = 2
                _ = fi27
                _ = v28
            }
        }
    }
    fmt.Println(v3, v4, v4)
    v29 := GenS1{F0: v4, F1: -128}
    var v30 int32 = -897
    if (154 == 34) {
        var v31 int32 = 0
        var v32 string = ("\n\t" + "")
        _ = v31
        _ = v32
    } else {
        v33 := []int{}
        _ = v33
    }
    var v34 bool = (v30 >= (^(v30)))
    v30 -= 2147483647
    v35 := GenS2{F0: (v4 - (v4 - v4)), F1: (v4 & (v4 & 32767)), F2: 978}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
}
