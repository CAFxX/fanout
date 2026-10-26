package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint8 = 127
    var v2 int16 = -535
    var v3 int = -2147483648
    var v4 uint = 65535
    var v5 uint16 = 961
    var v6 uint8 = (^(v1))
    var v7 int = (v3 / (((v3 / ((832) | 1))) | 1))
    fmt.Println(v6, v5, v1)
    var v8 bool = (!(!("" < "\n\t")))
    if true {
        var v9 float32 = 0.1
        var v10 uint8 = 33
        if true {
            v11 := GenS1{F0: 802}
            v2 -= (v2 - 0)
            v12 := GenS2{F0: v2, F1: 1}
            _ = v11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    if v8 {
        v13 := []int{}
        _ = v13
    }
    var v14 int32 = 41
    for fi15 := 0; fi15 < 3; fi15++ {
        v6 = (^(v6))
        _ = fi15
    }
    if v8 {
        v14 -= v14
        v16 := []int{}
        switch ((("hello" + "a") + ("a" + "a")) + "hello") {
        case "hello":
            v3 += 266
        case "\n\t":
            if false {
                var v17 uint64 = 650
                for fi18 := 0; fi18 < 9; fi18++ {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if v8 {
                            v5 = 26
                        }
                        switch "\n\t" {
                        case "":
                            v14 = v14
                        case "hello":
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        case "foo bar":
                            v20 := GenS3{F0: (v6 & (^(v6))), F1: 65535}
                            _ = v20
                        default:
                            switch -276 {
                            case -1:
                                var v21 int16 = -32768
                                _ = v21
                            case 4294967295:
                                fmt.Println(v3, v5, v17)
                            case 32767:
                                if len(v16) > 0 {
                                    v16[0] = 463
                                }
                            }
                        }
                        _ = fi19
                    }
                    v8 = ((((v5 ^ (^(v5))) < (v5 ^ (^(v5)))) && (930 != (v5 & v5))) && ((v14 < int32(v17)) || v8))
                    if false {
                        v22 := GenS1{F0: 65535}
                        for fi23 := 0; fi23 < 9; fi23++ {
                            fmt.Println(v7, v6)
                            _ = fi23
                        }
                        if v8 {
                            v8 = false
                        } else {
                            v7 = v7
                            _ = len(v16)
                            for fi24 := 0; fi24 < 4; fi24++ {
                                var v25 int8 = 127
                                fmt.Println(v14, v3)
                                _ = fi24
                                _ = v25
                            }
                        }
                        _ = v22
                    }
                    _ = fi18
                }
                v5 = 2
                _ = v17
            } else {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            }
        case "":
            fmt.Println(v6)
        default:
            v3 += v7
        }
        _ = v16
    } else {
        var v26 int64 = 385
        v27 := []int{}
        _ = v26
        _ = v27
    }
    fmt.Println(v4)
    var v28 float64 = -2.25
    v3 = v3
    v29 := GenS3{F0: (^(v6)), F1: 529}
    v5 -= 786
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS2{F0: (v2 + -1), F1: (v2 & v2)}
        v6 -= 2
        _ = fi30
        _ = v31
    }
    v1 -= 1
    v32 := []int{(v7 + v3)}
    fmt.Println(v6, v4, v6)
    fmt.Println(v5, v7, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
