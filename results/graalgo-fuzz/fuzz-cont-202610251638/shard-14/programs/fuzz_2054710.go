package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 uint32 }
type GenS4 struct { F0 int64; F1 bool; F2 int8 }
func genFn0(p0 uint64, p1 uint8) uint16 {
    v7 := GenS1{F0: -1, F1: 42}
    p0 = 4294967295
    v8 := GenS1{F0: 64, F1: 252}
    v9 := []int{9223372036854775807, -875, 32767}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint16 = 347
    var v2 uint64 = 932
    v3 := []uint32{0, 611}
    var v4 float64 = 1.5
    v5 := []string{}
    v6 := []int{9223372036854775807, 32767, 9223372036854775807}
    var v10 int8 = -92
    v5 = append(v5, "\n\t")
    var v11 string = ""
    for fi12 := 0; fi12 < 16; fi12++ {
        switch fi12 {
        case 127:
            v13 := GenS1{F0: (v10 << ((42) & 63)), F1: 952}
            _ = v13
        case -291:
            v2 += v2
        case 42:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            switch (("\n\t" + ("hello" + "hello")) + (("a" + "a") + (("" + ("hello" + "foo bar")) + ("\n\t" + v11)))) {
            case "":
                if (!true) {
                    fmt.Println(v1)
                    switch v11 {
                    case "":
                        fmt.Println(v1)
                    case "\n\t":
                        if true {
                            if true {
                                var v14 int16 = 1
                                fmt.Println(v4)
                                _ = v14
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v15 := GenS2{F0: (v1 % ((v1) | 1))}
                                v11 = (v11 + ((((v11 + v11) + "a") + "hello") + v11))
                                _ = v15
                            }
                        } else {
                            var v16 uint32 = 2147483648
                            var v17 uint32 = 826
                            _ = v16
                            _ = v17
                        }
                    }
                    fmt.Println(fi12)
                }
            case "hello":
                v18 := GenS3{F0: 1, F1: v10, F2: 0}
                _ = v18
            case "a":
                v19 := GenS4{F0: -9223372036854775808, F1: ((v1 ^ (v1 | v1)) <= 42), F2: v10}
                _ = v19
            }
        }
        fmt.Println(v2)
        _ = fi12
    }
    v1 = genFn0(18446744073709551615, 114)
    var v20 int16 = -1
    v20 += (v20 & v20)
    v21 := GenS4{F0: 1, F1: (!((!true) && false)), F2: 127}
    fmt.Println(v4)
    switch v11 {
    case "\n\t":
        var v22 float64 = (float64((float32(v4) / 1.0)) - v4)
        _ = v22
    case "a":
        fmt.Println(v1, v11, v10)
    case "hello":
        var v23 int64 = 0
        _ = v23
    }
    if (1 <= -495) {
        if true {
            var v24 string = (v11 + v11)
            _ = v24
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                if true {
                    var v26 uint64 = (v2 % ((v2) | 1))
                    if ((((v11 + "") + v11) + v11) == "a") {
                        if ((547 != 1) || true) {
                            var v27 int8 = v10
                            switch "\n\t" {
                            case "foo bar":
                                v20 = (v20 % ((-32768) | 1))
                            }
                            var v28 float64 = (float64(1.5) / (v4 + 1.5))
                            _ = v27
                            _ = v28
                        }
                    } else {
                        var v29 uint = 65535
                        v10 += v10
                        _ = v29
                    }
                    v30 := GenS4{F0: 1, F1: (int(v20) != fi25), F2: (v10 & -128)}
                    _ = v26
                    _ = v30
                }
                var v31 int8 = (v10 / ((-45) | 1))
                _ = fi25
                _ = v31
            }
            var v32 uint32 = 0
            _ = v32
        }
    } else {
        var v33 uint64 = 0
        _ = v33
    }
    var v34 uint16 = (^(v1))
    for fi35 := 0; fi35 < 20; fi35++ {
        for fi36 := 0; fi36 < 5; fi36++ {
            var v37 int64 = 42
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    if false {
        var v38 int8 = v10
        var v39 bool = false
        _ = v38
        _ = v39
    }
    fmt.Println(v1, v10, v11)
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
