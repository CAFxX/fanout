package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 string; F1 bool; F2 uint16 }

func main() {
    v1 := []string{"\n\t", ("" + "foo bar"), "\n\t", ""}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 float32 = 0.0
    _ = len(v1)
    v4 -= 3.14159
    v4 = v4
    if (false || true) {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 uint = 852
            _ = fi5
            _ = v6
        }
        if (!(0 < 42)) {
            for fi7 := 0; fi7 < 12; fi7++ {
                if false {
                    fmt.Println(v4)
                } else {
                    if true {
                        _ = len(v3)
                        var v8 uint32 = 615
                        _ = v8
                    } else {
                        _ = len(v1)
                        v9 := v2["a"]
                        fmt.Println(fi7)
                        _ = v9
                    }
                    v4 += (float32((3.14159 - float64((1e10 - (-2.25 - v4))))) - (float32(1.5) + v4))
                    var v10 uint8 = 44
                    _ = v10
                }
                if (-9223372036854775808 > 177) {
                    if ((!(false && true)) || (243 >= 2)) {
                        v11 := GenS1{F0: fi7}
                        if true {
                            if len(v1) > 0 {
                                v1[0] = ("foo bar" + "hello")
                            }
                        }
                        _ = v11
                    } else {
                        switch fi7 {
                        case -2147483648:
                            var v12 bool = (true || true)
                            _ = v12
                        }
                        v3[42] = fi7
                    }
                    var v13 float32 = v4
                    var v14 uint = 0
                    _ = v13
                    _ = v14
                }
                _ = fi7
            }
            var v15 int64 = 1
            _ = v15
        } else {
            var v16 string = "hello"
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 string = ""
                v18 = ((v16 + v16) + (("a" + (v16 + ((("\n\t" + v18) + (("foo bar" + v18) + "foo bar")) + ("\n\t" + "")))) + ("" + v18)))
                if (-2147483648 < -1) {
                    v16 = v16
                    var v19 uint = 191
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    }
    v4 -= 3.14159
    v20 := v2[(("a" + "") + ((("hello" + (("\n\t" + ("foo bar" + "")) + "hello")) + (("a" + ("" + "foo bar")) + "a")) + "foo bar"))]
    v20 -= v20
    for fi21 := 0; fi21 < 2; fi21++ {
        _ = len(v2)
        _ = fi21
    }
    if false {
        var v22 int32 = 1
        _ = v22
    } else {
        if (v4 <= v4) {
            v3[v20] = -2147483648
            if ((!true) && (0 > 0)) {
                if (true && true) {
                    fmt.Println(v20, v4)
                }
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    if (true && true) {
                        var v24 uint = 18446744073709551615
                        if true {
                            var v25 int16 = -32768
                            var v26 int16 = v25
                            _ = v25
                            _ = v26
                        }
                        _ = v24
                    }
                    _ = fi23
                }
                for fi27 := 0; fi27 < 9; fi27++ {
                    switch fi27 {
                    case -1:
                        v28 := GenS1{F0: 42}
                        _ = v28
                    case 128:
                        _ = len(v3)
                    case -2147483648:
                        switch ("a" + ("" + "a")) {
                        case "foo bar":
                            for fi29 := 0; fi29 < 8; fi29++ {
                                if false {
                                    v30 := GenS1{F0: fi27}
                                    if true {
                                        if false {
                                            v4 = 0.0
                                        } else {
                                            v2[("a" + ("foo bar" + "hello"))] = (v20 | (v20 << ((1) & 63)))
                                        }
                                        var v31 string = ""
                                        _ = v31
                                    }
                                    var v32 bool = true
                                    _ = v30
                                    _ = v32
                                }
                                _ = fi29
                            }
                        case "":
                            var v33 bool = (!false)
                            _ = v33
                        case "\n\t":
                            var v34 int = (^(fi27))
                            _ = v34
                        default:
                            if true {
                                fmt.Println(fi27, v4, v20)
                                if true {
                                    v20 = 471
                                }
                            }
                        }
                    }
                    _ = fi27
                }
                for fi35 := 0; fi35 < 17; fi35++ {
                    if len(v1) > 0 {
                        v1[0] = ("hello" + ((("hello" + "") + "a") + "a"))
                    }
                    switch fi35 {
                    case 930:
                        v36 := GenS2{F0: fi35, F1: 65535}
                        _ = v36
                    case -2147483648:
                        var v37 int16 = 42
                        _ = v37
                    }
                    delete(v3, 2)
                    _ = fi35
                }
            }
            var v38 int64 = -1
            _ = v38
        }
        v39 := GenS2{F0: 1, F1: 0}
        if false {
            switch "\n\t" {
            case "hello":
                var v40 string = "a"
                _ = v40
            case "":
                if (!(!("" >= "hello"))) {
                    v4 = v4
                } else {
                    v4 += -2.25
                    var v41 uint32 = (v20 | 42)
                    _ = v41
                }
            }
            switch "a" {
            case "foo bar":
                switch "foo bar" {
                case "hello":
                    var v42 int16 = 32767
                    _ = v42
                case "\n\t":
                    fmt.Println(v4, v4)
                case "foo bar":
                    for fi43 := 0; fi43 < 13; fi43++ {
                        var v44 int32 = 2147483647
                        if (!(-9223372036854775808 > 1)) {
                            v45 := GenS3{F0: "a", F1: false, F2: 371}
                            var v46 uint64 = 1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v45
                            _ = v46
                        }
                        if ((!false) || (795 >= (v44 >> ((65535) & 63)))) {
                            var v47 uint32 = (v20 + v20)
                            for fi48 := 0; fi48 < 4; fi48++ {
                                fmt.Println(v47, v44)
                                _ = fi48
                            }
                            _ = v47
                        } else {
                            var v49 float32 = v4
                            _ = v49
                        }
                        _ = fi43
                        _ = v44
                    }
                }
            }
        }
        _ = v39
    }
    delete(v2, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
