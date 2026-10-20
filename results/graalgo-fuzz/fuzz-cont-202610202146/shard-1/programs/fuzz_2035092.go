package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
type GenS6 struct { F0 bool; F1 int; F2 int }
func genFn0() string {
    var v7 uint = 2
    v8 := GenS1{F0: 42}
    var v9 uint64 = 42
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    v1 := []uint32{0}
    var v2 int32 = -573
    v3 := make(map[uint32]uint32)
    var v4 int = 40
    v5 := []uint32{1, 202, 2147483647}
    v6 := []string{"", "", "foo bar", "a"}
    v10 := v3[2147483647]
    v11 := GenS1{F0: 1}
    var v12 float32 = 3.14159
    var v13 uint = 0
    var v14 int64 = 0
    var v15 int8 = 0
    switch (-(v4)) {
    case 127:
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 int16 = int16(v15)
                v2 += 740
                _ = fi17
                _ = v18
            }
            var v19 uint32 = v10
            fmt.Println(v10, v10)
            _ = fi16
            _ = v19
        }
    case -9223372036854775808:
        v20 := GenS1{F0: 65535}
        _ = v20
    case 0:
        if true {
            fmt.Println(v2)
            var v21 string = "foo bar"
            if true {
                switch ((v21 + v21) + (("a" + "foo bar") + v21)) {
                case "hello":
                    switch (v4 >> ((v13) & 63)) {
                    case 602:
                        _ = len(v5)
                    }
                case "\n\t":
                    var v22 int8 = 1
                    _ = v22
                }
                if false {
                    v10 = v10
                    v23 := GenS1{F0: 32767}
                    _ = v23
                }
            }
            _ = v21
        }
    }
    var v24 uint64 = 152
    v12 += v12
    v25 := GenS2{F0: v24, F1: (v13 | 2)}
    if true {
        var v26 int32 = v2
        _ = v26
    } else {
        for fi27 := 0; fi27 < 14; fi27++ {
            _ = len(v5)
            v28 := GenS2{F0: v24, F1: v13}
            if true {
                for fi29 := 0; fi29 < 15; fi29++ {
                    v13 += (v13 % ((uint(v10)) | 1))
                    v15 = int8(fi27)
                    _ = len(v6)
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        v12 += v12
                    }
                    v13 = (v13 & (v13 - (v13 << ((201) & 63))))
                    v30 := GenS3{F0: (v24 + v24), F1: 232}
                    v24 = (v24 + v24)
                    _ = fi29
                    _ = v30
                }
            } else {
                if true {
                    v13 += (v13 - (v13 + (v13 & v13)))
                    var v31 float64 = 3.14159
                    _ = v31
                } else {
                    if false {
                        v2 -= v2
                        var v32 int32 = v2
                        if true {
                            if false {
                                if ((false || (!false)) || (false && ((int8(v10) <= v15) && (-1 < v32)))) {
                                    v33 := GenS4{F0: (v10 % ((v10) | 1))}
                                    v13 += (^(v13))
                                    _ = v33
                                } else {
                                    v34 := GenS2{F0: (^(v24)), F1: v13}
                                    switch int(v32) {
                                    case -2147483648:
                                        if true {
                                            v35 := GenS5{F0: (v2 ^ (^(v32))), F1: ((!false) && true), F2: true}
                                            _ = v35
                                        }
                                    case 2147483647:
                                        v15 += 0
                                    case 0:
                                        var v36 int16 = 1000
                                        _ = v36
                                    }
                                    for fi37 := 0; fi37 < 2; fi37++ {
                                        switch (^(v4)) {
                                        case -2147483648:
                                            v3[(v10 ^ v10)] = v10
                                            var v38 uint16 = 704
                                            _ = v38
                                        case -6:
                                            if (823 > uint16(v10)) {
                                                v15 = (v15 ^ (^(v15)))
                                            } else {
                                                var v39 uint = (^(v13))
                                                if true {
                                                    delete(v3, 1)
                                                    v5 = append(v5, (^(v10)))
                                                    fmt.Println(v15, v32, v15)
                                                }
                                                var v40 float32 = (v12 + v12)
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        fmt.Println(v12, v10, v14)
                                        if len(v5) > 0 {
                                            v5[0] = 981
                                        }
                                        _ = fi37
                                    }
                                    _ = v34
                                }
                            }
                        } else {
                            if ((32767 <= 370) || (((v2 << ((v13) & 63)) <= v32) && true)) {
                                fmt.Println(v2, v32)
                                var v41 int64 = (v14 ^ (v14 | -9223372036854775808))
                                _ = v41
                            } else {
                                v42 := GenS6{F0: true, F1: (v4 >> ((0) & 63)), F2: v4}
                                for fi43 := 0; fi43 < 10; fi43++ {
                                    v44 := GenS1{F0: 262}
                                    _ = fi43
                                    _ = v44
                                }
                                v45 := v3[65535]
                                _ = v42
                                _ = v45
                            }
                            var v46 int = fi27
                            v47 := GenS5{F0: v32, F1: false, F2: (!((((v32 | (v2 - (v2 << ((v13) & 63)))) > v2) || ((true && (v15 == v15)) || true)) || true))}
                            _ = v46
                            _ = v47
                        }
                        _ = v32
                    }
                    var v48 bool = true
                    _ = v48
                }
                v15 = v15
                var v49 uint8 = 2
                _ = v49
            }
            _ = fi27
            _ = v28
        }
        var v50 uint8 = 128
        _ = v50
    }
    var v51 uint32 = (v10 % ((v10) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
