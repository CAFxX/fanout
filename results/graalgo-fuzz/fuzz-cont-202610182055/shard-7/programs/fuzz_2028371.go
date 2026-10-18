package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int64; F1 uint }

func main() {
    v1 := []string{("hello" + ""), (("\n\t" + ("a" + ("hello" + "a"))) + "foo bar")}
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 bool = true
    var v5 uint16 = 65535
    for fi6 := 0; fi6 < 5; fi6++ {
        v4 = false
        _ = fi6
    }
    switch "hello" {
    case "":
        switch ("a" + ((((((("a" + "\n\t") + (("\n\t" + ("" + ("foo bar" + "foo bar"))) + "foo bar")) + ("\n\t" + "foo bar")) + ("" + ("" + ""))) + "\n\t") + "") + "\n\t")) {
        case "foo bar":
            for fi7 := 0; fi7 < 16; fi7++ {
                if false {
                    var v8 float64 = ((-2.25 - (-2.25 + 0.0)) * 0.0)
                    v4 = v4
                    _ = v8
                } else {
                    v4 = v4
                    fmt.Println(v4, fi7)
                }
                fmt.Println(v5, v5)
                _ = fi7
            }
        }
    }
    var v9 float64 = float64((float32((1.5 - 0.1)) - -2.25))
    switch ("a" + "") {
    case "hello":
        var v10 string = ""
        _ = v10
    }
    var v11 uint64 = 0
    if v4 {
        var v12 int16 = 1
        v5 = v5
        var v13 float64 = -2.25
        _ = v12
        _ = v13
    } else {
        var v14 string = "foo bar"
        _ = v14
    }
    if (v4 && (42 == -9223372036854775808)) {
        switch 32768 {
        case -2147483648:
            switch "a" {
            case "a":
                v11 += (v11 << ((4294967295) & 63))
            case "\n\t":
                _ = len(v1)
            }
        default:
            v3["\n\t"] = (("\n\t" + "foo bar") + "foo bar")
        }
    } else {
        if v4 {
            _ = len(v1)
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS1{F0: fi15, F1: 20}
                v17 := GenS2{F0: "foo bar"}
                if ((^(v5)) > (^(v5))) {
                    v18 := GenS2{F0: ((("" + ("\n\t" + "\n\t")) + "") + ("a" + "hello"))}
                    switch fi15 {
                    case 32768:
                        if ((!false) && (false && (!(uint32(v11) > 1)))) {
                            var v19 string = (("hello" + "hello") + "hello")
                            _ = v19
                        } else {
                            v5 = 255
                        }
                    }
                    switch (-(fi15)) {
                    case -622:
                        var v20 int64 = 9223372036854775807
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 12; fi21++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                _ = len(v1)
                                v22 := GenS2{F0: "a"}
                                _ = v22
                            } else {
                                if (-2147483648 <= 1) {
                                    var v23 uint8 = uint8(v11)
                                    v24 := GenS2{F0: ""}
                                    v9 = v9
                                    _ = v23
                                    _ = v24
                                }
                                if ((1.5 > 1.5) && (9223372036854775807 >= fi21)) {
                                    if false {
                                        v3[(("hello" + (("" + "hello") + ("a" + "\n\t"))) + (("a" + ("" + ("hello" + ""))) + "foo bar"))] = "\n\t"
                                    } else {
                                        if v4 {
                                            v3[(("" + "foo bar") + "foo bar")] = ((((("hello" + ("a" + ("a" + "hello"))) + "a") + "hello") + "foo bar") + ("a" + ("\n\t" + (((("a" + "") + "") + (("hello" + ("hello" + (("\n\t" + "hello") + "hello"))) + "")) + "\n\t"))))
                                            v5 -= uint16(v11)
                                            v25 := GenS1{F0: (fi21 * (fi15 | 9223372036854775807)), F1: 65535}
                                            _ = v25
                                        } else {
                                            var v26 int32 = 1
                                            v26 = 245
                                            v3["foo bar"] = "hello"
                                            _ = v26
                                        }
                                        if true {
                                            v27 := GenS2{F0: "hello"}
                                            _ = v27
                                        }
                                        fmt.Println(v9, fi21)
                                    }
                                    var v28 bool = (!(v11 <= (^(v11))))
                                    var v29 uint32 = 2
                                    _ = v28
                                    _ = v29
                                }
                            }
                            _ = fi21
                        }
                    }
                    _ = v18
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            v5 = (v5 * (v5 & v5))
        } else {
            var v30 uint64 = v11
            switch "" {
            case "foo bar":
                v31 := GenS1{F0: -1, F1: 1}
                _ = v31
            case "hello":
                switch 4294967295 {
                case -236:
                    var v32 float32 = -2.25
                    _ = v32
                case 32767:
                    for fi33 := 0; fi33 < 7; fi33++ {
                        v34 := GenS2{F0: "\n\t"}
                        _ = fi33
                        _ = v34
                    }
                case -499:
                    var v35 uint16 = (v5 & v5)
                    _ = v35
                }
            case "\n\t":
                var v36 uint64 = (v30 << ((0) & 63))
                _ = v36
            }
            _ = v30
        }
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    v37 := GenS3{F0: 791, F1: (!(!(42 != 2)))}
    v38 := GenS4{F0: 0, F1: 318}
    fmt.Println(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v9
}
