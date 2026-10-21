package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 int }

func main() {
    var v1 uint64 = 2
    v2 := []string{"hello"}
    var v3 float32 = (1.5 * 1.5)
    switch 826 {
    case 9223372036854775807:
        v3 = v3
    case 32767:
        for fi4 := 0; fi4 < 16; fi4++ {
            fmt.Println(fi4, v1)
            if false {
                for fi5 := 0; fi5 < 9; fi5++ {
                    var v6 uint16 = 2
                    _ = fi5
                    _ = v6
                }
                var v7 int = (-(fi4))
                _ = v7
            }
            fmt.Println(fi4)
            _ = fi4
        }
    case -2147483648:
        v8 := GenS1{F0: ((168 < 2147483647) || (false && (!true))), F1: 257}
        _ = v8
    default:
        var v9 string = "a"
        _ = v9
    }
    switch -1 {
    case 32768:
        v1 = (v1 % (((^(v1))) | 1))
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if false {
            v10 := GenS2{F0: 74, F1: -203, F2: (v3 <= 3.14159)}
            for fi11 := 0; fi11 < 15; fi11++ {
                fmt.Println(fi11, fi11, fi11)
                var v12 float64 = float64(3.14159)
                var v13 int16 = -32768
                _ = fi11
                _ = v12
                _ = v13
            }
            v1 += (v1 - v1)
            _ = v10
        } else {
            fmt.Println(v1, v3)
            if (832 >= 1) {
                var v14 uint = 42
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 string = ""
                        fmt.Println(fi15, fi15, v3)
                        _ = fi16
                        _ = v17
                    }
                    var v18 int16 = -1
                    _ = fi15
                    _ = v18
                }
                _ = v14
            }
            v2 = append(v2, ("hello" + "foo bar"))
        }
    }
    fmt.Println(v1)
    v19 := GenS2{F0: -384, F1: -9223372036854775808, F2: (!(false || (625 > 1)))}
    if true {
        switch 127 {
        case 128:
            for fi20 := 0; fi20 < 11; fi20++ {
                v21 := GenS3{F0: 1000, F1: 0}
                fmt.Println(v1)
                _ = fi20
                _ = v21
            }
        case 32768:
            if false {
                switch 42 {
                case 588:
                    var v22 int16 = 1
                    _ = v22
                case 4294967295:
                    switch 6 {
                    case -2147483648:
                        for fi23 := 0; fi23 < 11; fi23++ {
                            for fi24 := 0; fi24 < 4; fi24++ {
                                switch fi23 {
                                case 9223372036854775807:
                                    var v25 int32 = 0
                                    _ = v25
                                case 128:
                                    v2 = append(v2, "\n\t")
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi24
                            }
                            var v26 int16 = -32768
                            for fi27 := 0; fi27 < 19; fi27++ {
                                var v28 bool = true
                                _ = fi27
                                _ = v28
                            }
                            _ = fi23
                            _ = v26
                        }
                    }
                case 42:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v29 string = "foo bar"
                _ = v29
            }
        case 127:
            for fi30 := 0; fi30 < 10; fi30++ {
                v2 = append(v2, (("foo bar" + ("hello" + "a")) + "a"))
                for fi31 := 0; fi31 < 7; fi31++ {
                    v32 := GenS3{F0: 32, F1: fi30}
                    v3 = v3
                    _ = fi31
                    _ = v32
                }
                for fi33 := 0; fi33 < 10; fi33++ {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = fi33
                }
                _ = fi30
            }
        }
        v1 -= v1
    }
    var v34 uint = 4294967295
    v34 = (v34 | 1)
    var v35 int = -2147483648
    var v36 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
