package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8, p1 int, p2 int16) int8 {
    var v6 uint32 = 2147483648
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 -= p0
            p2 = (^(p2))
            p0 = p0
            switch (("hello" + ("foo bar" + "")) + "foo bar") {
            case "a":
                if false {
                    if false {
                        v9 := []int{(p1 - 9223372036854775807), (^(p1)), fi8, 195}
                        fmt.Println(fi7)
                        _ = v9
                    }
                    v10 := GenS1{F0: uint16(v6)}
                    v6 = (v6 & 2147483647)
                    _ = v10
                } else {
                    var v11 uint = 0
                    _ = v11
                }
            case "\n\t":
                var v12 int16 = -987
                _ = v12
            case "foo bar":
                p1 += 128
            }
            _ = fi8
        }
        _ = fi7
    }
    if false {
        v13 := []int{(p1 & 32768)}
        _ = v13
    } else {
        v6 += (v6 / ((v6) | 1))
        switch "\n\t" {
        case "hello":
            var v14 float64 = -2.25
            _ = v14
        }
        fmt.Println(p2, p0)
    }
    if true {
        v15 := []int{(p1 / (((^(p1))) | 1))}
        _ = v15
    }
    if ((false && (!true)) && true) {
        fmt.Println(p0)
        var v16 uint = 65535
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{128}
    v3 := []uint32{42}
    var v4 uint = 4294967295
    var v5 uint16 = 1
    for fi17 := 0; fi17 < 15; fi17++ {
        if false {
            switch ("hello" + ("hello" + "foo bar")) {
            case "hello":
                for fi18 := 0; fi18 < 5; fi18++ {
                    v4 = 1
                    v1 -= ((((((v1 + (v1 - 3.14159)) * v1) - (v1 / 0.1)) + v1) + v1) + v1)
                    _ = fi18
                }
            case "":
                switch fi17 {
                case 127:
                    v4 -= 2
                case 128:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        if true {
                            v5 += v5
                            var v20 int8 = 42
                            var v21 int64 = 42
                            v21 += (v21 << ((991) & 63))
                            var v22 uint16 = (v5 / ((674) | 1))
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS1{F0: v5}
                            if (v5 != v5) {
                                v5 += (^(v5))
                                v4 = 42
                            } else {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v1 += float32(0.0)
                                    v24 := GenS1{F0: v5}
                                    _ = v24
                                }
                                v4 = (v4 & (v4 << (((^(v4))) & 63)))
                            }
                            var v25 int8 = 121
                            _ = v23
                            _ = v25
                        }
                        _ = fi19
                    }
                }
            case "\n\t":
                v5 -= (v5 ^ v5)
            }
        }
        switch fi17 {
        case 0:
            v1 = (v1 - -2.25)
        default:
            v1 = 0.1
        }
        v26 := GenS1{F0: v5}
        _ = fi17
        _ = v26
    }
    v2 = append(v2, 128)
    for fi27 := 0; fi27 < 3; fi27++ {
        if (((1 > (v5 / (((v5 & (^(v5)))) | 1))) && (genFn0(1, int(v4), 0) == -23)) && true) {
            if (-841 == 32767) {
                var v28 float64 = 1.5
                _ = v28
            }
            v3 = append(v3, 42)
            for fi29 := 0; fi29 < 6; fi29++ {
                v4 = v4
                v5 = v5
                if false {
                    v4 -= v4
                    if (v4 >= v4) {
                        var v30 uint32 = 42
                        var v31 string = "hello"
                        v32 := GenS2{F0: 2147483647}
                        _ = v30
                        _ = v31
                        _ = v32
                    } else {
                        v4 = v4
                    }
                    v33 := GenS1{F0: v5}
                    _ = v33
                }
                _ = fi29
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi27
    }
    v34 := GenS2{F0: -1}
    v3 = append(v3, 1)
    if (!(1 < 0)) {
        var v35 int8 = 22
        _ = v35
    } else {
        var v36 uint64 = 0
        _ = v36
    }
    if true {
        fmt.Println(v4)
    } else {
        v1 -= float32(float64(-2.25))
    }
    v1 = 0.0
    fmt.Println(v4, v4)
    fmt.Println(v1, v5)
    for fi37 := 0; fi37 < 11; fi37++ {
        for fi38 := 0; fi38 < 3; fi38++ {
            v39 := GenS3{F0: 0}
            _ = len(v2)
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    v40 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v5
}
