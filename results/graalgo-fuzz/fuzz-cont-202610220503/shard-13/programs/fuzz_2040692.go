package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 string = "\n\t"
    var v2 int32 = 29
    var v3 uint = 18446744073709551615
    var v4 float64 = -2.25
    v4 = v4
    v5 := []int{715, 415}
    v1 = "\n\t"
    v1 = ((v1 + v1) + "\n\t")
    if false {
        v6 := GenS1{F0: false, F1: -128, F2: 1}
        fmt.Println(v1, v3)
        switch v1 {
        case "\n\t":
            var v7 uint16 = 135
            _ = v7
        case "":
            fmt.Println(v4, v4, v2)
        default:
            if true {
                if (v2 != v2) {
                    if ((-1 < -9223372036854775808) || (int32(v3) >= v2)) {
                        if false {
                            var v8 uint32 = 768
                            if len(v5) > 0 {
                                v5[0] = -698
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v8
                        }
                        if false {
                            v9 := GenS1{F0: true, F1: 0, F2: 1}
                            v2 -= v2
                            _ = v9
                        } else {
                            switch -9223372036854775808 {
                            case 0:
                                v4 += (v4 + (1.5 - ((v4 - v4) - float64(3.14159))))
                            }
                        }
                    }
                    if false {
                        v2 += (v2 | v2)
                    } else {
                        for fi10 := 0; fi10 < 12; fi10++ {
                            v11 := GenS2{F0: 42}
                            _ = len(v5)
                            _ = fi10
                            _ = v11
                        }
                        v12 := GenS2{F0: -32768}
                        _ = v12
                    }
                }
            } else {
                v2 = -2147483648
                switch 311 {
                case 32767:
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 uint64 = 18446744073709551615
                        fmt.Println(v3, v1)
                        switch (^(fi13)) {
                        case -1:
                            var v15 uint16 = 0
                            _ = v15
                        case -2147483648:
                            for fi16 := 0; fi16 < 13; fi16++ {
                                if false {
                                    _ = len(v5)
                                } else {
                                    switch fi13 {
                                    case 128:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                                _ = fi16
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                }
                var v17 uint16 = 571
                _ = v17
            }
        }
        _ = v6
    }
    var v18 int = 1
    var v19 int16 = 32767
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int64 = 9223372036854775807
        for fi22 := 0; fi22 < 2; fi22++ {
            _ = len(v5)
            v18 = 9223372036854775807
            var v23 int16 = v19
            _ = fi22
            _ = v23
        }
        switch 32767 {
        case -550:
            v18 = 9223372036854775807
        }
        _ = fi20
        _ = v21
    }
    if (2147483647 == 2147483647) {
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 bool = (uint64(v3) > 760)
            var v26 int16 = 1
            v2 = 2147483647
            _ = fi24
            _ = v25
            _ = v26
        }
        v27 := GenS1{F0: (false || false), F1: 1, F2: 127}
        var v28 uint64 = 42
        _ = v27
        _ = v28
    } else {
        var v29 uint64 = 396
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := GenS3{F0: v19}
            if (!(v18 >= fi31)) {
                if false {
                    var v33 int64 = -1
                    _ = v33
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v34 string = (v1 + ((("foo bar" + ("\n\t" + v1)) + (v1 + ((v1 + "hello") + "foo bar"))) + v1))
                _ = v34
            }
            fmt.Println(v3, v18)
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v19 += (v19 ^ v19)
    var v35 uint8 = 150
    var v36 bool = true
    if (v19 != (v19 * (v19 + v19))) {
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
    }
    var v37 uint = v3
    var v38 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
