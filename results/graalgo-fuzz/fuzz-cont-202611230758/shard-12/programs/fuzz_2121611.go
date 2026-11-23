package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int8; F2 bool }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int; F1 uint8; F2 string }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 int32 }
func genFn0(p0 uint64, p1 uint64, p2 uint32) uint64 {
    if false {
        var v7 bool = (1 <= 1)
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 255}
    var v9 uint16 = 65535
    p0 = p1
    v9 += (v9 >> ((876) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2147483648
    v3 := make(map[int]int)
    var v4 int16 = -182
    var v5 int16 = (v4 % ((v4) | 1))
    var v6 int32 = 0
    var v10 int = 971
    v5 = -651
    for fi11 := 0; fi11 < 20; fi11++ {
        v4 = 820
        _ = fi11
    }
    fmt.Println(v10, v5, v5)
    for fi12 := 0; fi12 < 2; fi12++ {
        switch 2147483647 {
        case -2147483648:
            v5 += (v5 ^ v4)
        case 0:
            v13 := GenS1{F0: 0, F1: 128}
            _ = v13
        case 4294967295:
            if (9223372036854775807 < 0) {
                var v14 uint64 = 904
                v15 := GenS2{F0: "a", F1: 2, F2: fi12}
                if (42 <= v2) {
                    var v16 string = "hello"
                    var v17 int64 = -1
                    _ = v16
                    _ = v17
                }
                _ = v14
                _ = v15
            } else {
                if false {
                    v18 := GenS3{F0: "foo bar", F1: 122}
                    v19 := []int{(v10 | (fi12 << ((4294967295) & 63)))}
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS4{F0: (1 != 42)}
                    _ = v20
                }
            }
        default:
            var v21 uint32 = v2
            _ = v21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint = 216
            _ = fi22
            _ = v23
        }
        if false {
            v24 := GenS1{F0: 42, F1: 2}
            switch "\n\t" {
            case "foo bar":
                delete(v3, 127)
            case "\n\t":
                var v25 float64 = -2.25
                _ = v25
            case "hello":
                v26 := GenS5{F0: 406, F1: 100, F2: false}
                _ = v26
            }
            var v27 uint8 = 127
            _ = v24
            _ = v27
        }
        _ = fi12
    }
    v5 += 32767
    v10 -= -198
    v2 += 2
    v28 := GenS3{F0: (("a" + ("hello" + "\n\t")) + ""), F1: 100}
    var v29 int8 = 100
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 int8 = 2
        var v32 float64 = 0.0
        _ = fi30
        _ = v31
        _ = v32
    }
    v4 += v4
    v33 := GenS2{F0: (("a" + "foo bar") + ("a" + "foo bar")), F1: 167, F2: (v10 / ((42) | 1))}
    var v34 int16 = v4
    var v35 int64 = 1
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := []int{(fi36 + v10)}
        for fi38 := 0; fi38 < 10; fi38++ {
            if false {
                switch 32768 {
                case -1:
                    if ((v35 + 42) <= v35) {
                        fmt.Println(v4, v34)
                        if len(v37) >= 2 {
                            v37 = v37[1:len(v37)]
                        }
                        var v39 int16 = v5
                        _ = v39
                    }
                case -2147483648:
                    v29 -= (v29 % ((-128) | 1))
                case 42:
                    fmt.Println(v5, v34)
                    v29 += -1
                    var v40 bool = (!((false || false) && true))
                    _ = v40
                default:
                    if true {
                        v2 = v2
                    } else {
                        var v41 bool = true
                        var v42 uint16 = 65535
                        v43 := GenS6{F0: 2, F1: (v4 | 32767)}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
            }
            if false {
                switch ((("foo bar" + "hello") + "hello") + "foo bar") {
                case "a":
                    v44 := GenS7{F0: (v10 - int(v4)), F1: 44, F2: "a"}
                    _ = v44
                case "":
                    var v45 float32 = (-2.25 * 0.1)
                    _ = v45
                default:
                    fmt.Println(v5)
                }
                switch 42 {
                case 1:
                    switch (-(fi36)) {
                    case 412:
                        if false {
                            switch 9223372036854775807 {
                            case -138:
                                fmt.Println(v2, v10, v6)
                            case -50:
                                var v46 int64 = (v35 >> ((706) & 63))
                                _ = v46
                            }
                            var v47 uint = 42
                            _ = v47
                        } else {
                            var v48 uint8 = 127
                            switch int(v48) {
                            case 32767:
                                fmt.Println(fi38, v48)
                            case -887:
                                v4 = (v4 ^ -1)
                            }
                            fmt.Println(v29, fi36, fi36)
                            _ = v48
                        }
                    case 32767:
                        switch "" {
                        case "hello":
                            fmt.Println(fi36, v10)
                        case "a":
                            v49 := GenS3{F0: (("hello" + ("\n\t" + "\n\t")) + "\n\t"), F1: -1}
                            _ = v49
                        }
                    case 0:
                        var v50 uint = 1
                        _ = v50
                    }
                case -9223372036854775808:
                    switch "" {
                    case "foo bar":
                        var v51 int = (fi38 / (((fi36 / ((fi36) | 1))) | 1))
                        _ = v51
                    case "":
                        if (0.1 > 3.14159) {
                            v52 := GenS8{F0: (fi38 % ((fi36) | 1)), F1: 42}
                            var v53 int64 = -1
                            delete(v3, (^(v10)))
                            _ = v52
                            _ = v53
                        }
                    case "\n\t":
                        var v54 int8 = v29
                        _ = v54
                    }
                default:
                    var v55 bool = ("hello" > "\n\t")
                    v4 = v4
                    _ = v55
                }
            }
            _ = fi38
        }
        _ = fi36
        _ = v37
    }
    v56 := GenS9{F0: v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
