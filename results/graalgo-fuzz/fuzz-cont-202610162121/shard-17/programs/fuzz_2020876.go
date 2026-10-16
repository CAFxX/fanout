package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    v3 := []uint32{1, 42}
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 997}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        switch "" {
        case "":
            if false {
                var v7 int16 = 0
                v7 = v7
                var v8 string = "foo bar"
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch "foo bar" {
                    case "hello":
                        _ = len(v2)
                    case "\n\t":
                        var v10 uint = 65535
                        _ = v10
                    }
                    v11 := GenS2{F0: -1}
                    _ = fi9
                    _ = v11
                }
            }
        case "a":
            var v12 bool = true
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 1; fi13++ {
                switch "" {
                case "\n\t":
                    fmt.Println(fi13)
                default:
                    v2[(("hello" + "\n\t") + ("a" + "hello"))] = ("hello" + "foo bar")
                }
                _ = fi13
            }
        default:
            var v14 float64 = 1e10
            _ = v14
        }
        fmt.Println(fi6)
        _ = fi6
    }
    var v15 uint = 255
    var v16 uint64 = 190
    fmt.Println(v16, v15)
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            v18 := GenS2{F0: 1}
            if (!((127 != 97) && (811 == -1))) {
                for fi19 := 0; fi19 < 8; fi19++ {
                    if (474 >= -682) {
                        var v20 string = ""
                        v15 = v15
                        _ = v20
                    } else {
                        var v21 uint64 = 18446744073709551615
                        if true {
                            v2[""] = (("foo bar" + ("\n\t" + ("a" + ""))) + ("foo bar" + ((("foo bar" + ("\n\t" + ("" + "a"))) + ((("a" + "") + "foo bar") + (("a" + "hello") + "\n\t"))) + "\n\t")))
                            fmt.Println(fi17, fi19, fi19)
                        }
                        _ = v21
                    }
                    v22 := GenS1{F0: 525}
                    _ = fi19
                    _ = v22
                }
                if true {
                    if false {
                        v15 -= (v15 + 18446744073709551615)
                        v15 = v15
                    }
                    v23 := GenS1{F0: 2}
                    _ = v23
                } else {
                    switch "" {
                    case "hello":
                        for fi24 := 0; fi24 < 2; fi24++ {
                            var v25 uint32 = 4294967295
                            v26 := GenS3{F0: 2, F1: 42}
                            var v27 int8 = 96
                            if (!("hello" == "")) {
                                if false {
                                    fmt.Println(v16, v16)
                                } else {
                                    var v28 int = (fi24 ^ fi17)
                                    v29 := GenS1{F0: 255}
                                    var v30 uint16 = 0
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case "a":
                        var v31 float32 = 1.5
                        _ = v31
                    }
                }
                fmt.Println(v15)
            } else {
                _ = len(v3)
                v32 := GenS1{F0: 0}
                var v33 uint64 = (v16 << ((v15) & 63))
                _ = v32
                _ = v33
            }
            if true {
                var v34 string = "foo bar"
                _ = v34
            }
            _ = v18
        } else {
            var v35 int = fi17
            _ = v35
        }
        if false {
            if ("hello" <= ("foo bar" + "")) {
                for fi36 := 0; fi36 < 16; fi36++ {
                    v15 -= v15
                    var v37 int64 = int64(v16)
                    v38 := v2[""]
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v15 -= (v15 - v15)
                for fi39 := 0; fi39 < 2; fi39++ {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    _ = fi39
                }
            }
            var v40 int16 = 42
            var v41 float32 = float32(3.14159)
            _ = v40
            _ = v41
        } else {
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                if true {
                    v42 := GenS3{F0: 255, F1: -103}
                    if (-1 >= -1) {
                        fmt.Println(v15, v15, v15)
                        fmt.Println(fi17, v16)
                    }
                    v16 = 0
                    _ = v42
                }
            }
            fmt.Println(v15, v16, v16)
            v16 -= v16
        }
        switch fi17 {
        case 684:
            var v43 uint16 = 2
            _ = v43
        case 32768:
            v16 += v16
        default:
            if (-766 != 42) {
                var v44 bool = false
                if true {
                    v15 = v15
                    v44 = (125 == 42)
                    for fi45 := 0; fi45 < 4; fi45++ {
                        v46 := GenS3{F0: 165, F1: 0}
                        _ = fi45
                        _ = v46
                    }
                } else {
                    switch int(v16) {
                    case 881:
                        var v47 uint64 = 1
                        _ = v47
                    case -9223372036854775808:
                        var v48 int8 = 1
                        _ = v48
                    case 2147483647:
                        v44 = (!(9223372036854775807 <= -1))
                    }
                }
                v2[(("a" + (("\n\t" + "\n\t") + ("a" + "hello"))) + "hello")] = ((("foo bar" + (("\n\t" + "a") + (("hello" + "a") + "a"))) + "\n\t") + ("a" + (("" + "hello") + ("a" + "\n\t"))))
                _ = v44
            } else {
                _ = len(v2)
                for fi49 := 0; fi49 < 12; fi49++ {
                    if ((v16 + 1) == 100) {
                        var v50 uint8 = 172
                        v16 = v16
                        _ = v50
                    } else {
                        v16 -= v16
                    }
                    var v51 int32 = 42
                    _ = fi49
                    _ = v51
                }
            }
        }
        _ = fi17
    }
    v15 = (^(v15))
    switch -1 {
    case -437:
        v15 = 42
    case 32768:
        switch 127 {
        case -359:
            v16 = v16
        case 1:
            v52 := GenS4{F0: 45, F1: 203, F2: -651}
            _ = v52
        }
    default:
        v16 = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
