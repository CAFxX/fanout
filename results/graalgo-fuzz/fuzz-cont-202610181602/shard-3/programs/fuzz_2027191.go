package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        switch (fi4 / (((fi4 - fi4)) | 1)) {
        case -554:
            v5 := []int{fi4, fi4, fi4}
            _ = v5
        case 0:
            if (!false) {
                v6 := GenS1{F0: "\n\t"}
                var v7 uint = 18446744073709551615
                var v8 uint8 = 157
                _ = v6
                _ = v7
                _ = v8
            }
        case -761:
            var v9 bool = false
            _ = v9
        }
        _ = fi4
    }
    v10 := []int{}
    _ = v10
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 float64 = (float64(-2.25) + 1.5)
    v3 -= (0.1 * ((-2.25 - v3) + v3))
    for fi11 := 0; fi11 < 12; fi11++ {
        if (!((837 >= 0) || ((((42 <= 32767) && (true || false)) || (true || (-9223372036854775808 != -9223372036854775808))) && false))) {
            _ = len(v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v12 int64 = 1
            _ = v12
        }
        _ = fi11
    }
    v13 := GenS1{F0: (("hello" + "foo bar") + "a")}
    v3 = 0.1
    var v14 int8 = 100
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint8 = 127
            switch "hello" {
            case "\n\t":
                _ = len(v1)
            default:
                switch ("\n\t" + "foo bar") {
                case "a":
                    v17 = 255
                case "hello":
                    v18 := GenS1{F0: (("a" + ((("hello" + "a") + "a") + "")) + "")}
                    _ = v18
                }
            }
            if ((("" + "a") + "hello") <= "\n\t") {
                _ = len(v1)
            } else {
                if true {
                    switch fi16 {
                    case 4294967295:
                        v19 := GenS1{F0: ("foo bar" + "a")}
                        _ = v19
                    case 2147483647:
                        var v20 int8 = v14
                        _ = v20
                    default:
                        switch "hello" {
                        case "\n\t":
                            var v21 int8 = 1
                            _ = v21
                        case "foo bar":
                            var v22 int32 = -1
                            _ = v22
                        default:
                            if false {
                                if true {
                                    v23 := GenS1{F0: "a"}
                                    var v24 uint = 1
                                    _ = v23
                                    _ = v24
                                } else {
                                    delete(v1, "\n\t")
                                }
                            }
                        }
                        switch "\n\t" {
                        case "a":
                            v25 := GenS1{F0: (("a" + "foo bar") + "")}
                            _ = v25
                        case "foo bar":
                            v26 := GenS2{F0: fi15}
                            _ = v26
                        default:
                            v27 := GenS3{F0: (-(v14)), F1: "foo bar"}
                            _ = v27
                        }
                        switch fi16 {
                        case 42:
                            if false {
                                v28 := GenS2{F0: (fi15 % ((-1) | 1))}
                                if (!((0.1 > 0.1) || (("foo bar" >= "hello") || (false || (-2147483648 != -1))))) {
                                    var v29 int64 = -1
                                    _ = v29
                                }
                                _ = v28
                            }
                        }
                    }
                }
            }
            _ = fi16
            _ = v17
        }
        v30 := GenS4{F0: 2, F1: ("hello" + "\n\t")}
        _ = len(v2)
        _ = fi15
        _ = v30
    }
    if (1 > -509) {
        if true {
            v31 := GenS5{F0: v14}
            var v32 uint = 1
            v3 = v3
            _ = v31
            _ = v32
        }
        if false {
            v14 = v14
            var v33 uint64 = 2
            _ = v33
        } else {
            var v34 uint32 = 2147483647
            v35 := GenS2{F0: 0}
            _ = len(v1)
            _ = v34
            _ = v35
        }
        var v36 int16 = 0
        _ = v36
    }
    v2 = append(v2, "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
