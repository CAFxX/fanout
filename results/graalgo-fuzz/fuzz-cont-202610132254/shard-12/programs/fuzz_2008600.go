package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 string) uint32 {
    var v6 uint64 = 2
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: (^(v6)), F1: -128}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 953
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]uint32)
    var v4 string = "\n\t"
    var v5 uint = 4294967295
    v5 = v5
    v9 := GenS1{F0: 4294967295, F1: -75}
    fmt.Println(v5)
    v10 := GenS1{F0: 4294967295, F1: 127}
    for fi11 := 0; fi11 < 20; fi11++ {
        v5 += 59
        if false {
            var v12 bool = (fi11 <= (fi11 ^ (fi11 * (fi11 / ((-297) | 1)))))
            var v13 int8 = 31
            _ = v12
            _ = v13
        } else {
            var v14 int16 = 0
            _ = v14
        }
        delete(v1, ("foo bar" + v4))
        _ = fi11
    }
    if (934 == 65535) {
        if (v5 >= (v5 * v5)) {
            delete(v3, 32767)
            v15 := GenS1{F0: 390, F1: 63}
            var v16 string = "hello"
            _ = v15
            _ = v16
        }
        var v17 uint = v5
        _ = v17
    } else {
        v5 += v5
        if (!("\n\t" >= v4)) {
            if false {
                var v18 float64 = 0.1
                v19 := []int{4294967295, 2147483647}
                var v20 uint64 = 18446744073709551615
                _ = v18
                _ = v19
                _ = v20
            } else {
                v21 := []int{}
                var v22 int32 = 522
                switch 4294967295 {
                case 0:
                    v21 = append(v21, -471)
                case 824:
                    if true {
                        fmt.Println(v2)
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        var v23 int8 = -47
                        _ = v23
                    }
                }
                _ = v21
                _ = v22
            }
            fmt.Println(v2)
        }
        if false {
            v24 := []int{0, 32767, 9223372036854775807, -47}
            var v25 int64 = v2
            if (0 < -1) {
                if (v4 == "") {
                    var v26 uint8 = 52
                    var v27 int8 = 0
                    var v28 float32 = float32(1.5)
                    _ = v26
                    _ = v27
                    _ = v28
                }
                v29 := GenS1{F0: 4294967295, F1: 42}
                v2 -= v2
                _ = v29
            }
            _ = v24
            _ = v25
        } else {
            if false {
                var v30 string = ""
                v31 := GenS2{F0: v5, F1: 1}
                var v32 uint16 = 255
                _ = v30
                _ = v31
                _ = v32
            } else {
                if (-2.25 != (float32((3.14159 * ((3.14159 + (1.5 - (1.5 / 3.14159))) - (float64(0.0) / 0.1)))) / (0.1 / (1.5 * -2.25)))) {
                    switch v4 {
                    case "foo bar":
                        v33 := GenS3{F0: 0}
                        _ = v33
                    case "a":
                        if ((v5 + 42) < (^(v5))) {
                            switch -984 {
                            case -2147483648:
                                if false {
                                    var v34 uint32 = 1
                                    _ = v34
                                } else {
                                    v35 := GenS1{F0: 1, F1: -99}
                                    var v36 uint32 = 42
                                    _ = v35
                                    _ = v36
                                }
                            case -338:
                                v37 := GenS3{F0: 0}
                                _ = v37
                            case 756:
                                delete(v1, v4)
                            default:
                                for fi38 := 0; fi38 < 1; fi38++ {
                                    var v39 int32 = 30
                                    v2 -= v2
                                    v40 := GenS2{F0: v5, F1: 560}
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                        } else {
                            fmt.Println(v5)
                            var v41 int = 127
                            fmt.Println(v41, v4)
                            _ = v41
                        }
                    default:
                        fmt.Println(v4)
                    }
                    for fi42 := 0; fi42 < 3; fi42++ {
                        var v43 bool = (!true)
                        v3[(^(fi42))] = genFn0(("hello" + v4))
                        for fi44 := 0; fi44 < 2; fi44++ {
                            var v45 int32 = 2147483647
                            _ = fi44
                            _ = v45
                        }
                        _ = fi42
                        _ = v43
                    }
                } else {
                    for fi46 := 0; fi46 < 13; fi46++ {
                        fmt.Println(v5)
                        _ = fi46
                    }
                }
                v2 += (v2 / ((-9223372036854775808) | 1))
                v47 := GenS1{F0: 2, F1: -1}
                _ = v47
            }
            var v48 uint16 = 222
            _ = v48
        }
    }
    var v49 string = ""
    for fi50 := 0; fi50 < 20; fi50++ {
        v51 := GenS1{F0: 18446744073709551615, F1: -1}
        if true {
            for fi52 := 0; fi52 < 4; fi52++ {
                v4 = v4
                fmt.Println(fi50, v2, fi52)
                _ = len(v3)
                _ = fi52
            }
        }
        _ = fi50
        _ = v51
    }
    v49 = ((v49 + v49) + ("hello" + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v9
}
