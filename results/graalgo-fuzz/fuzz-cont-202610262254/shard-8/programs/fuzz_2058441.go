package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int64 }

func main() {
    var v1 uint16 = 306
    v2 := []string{"hello", ""}
    v3 := make(map[string]uint32)
    var v4 uint16 = (v1 - 65535)
    v5 := []uint32{81}
    var v6 float64 = 3.14159
    var v7 bool = ((0 <= 59) && (true || (!false)))
    v4 = (v4 * 65535)
    var v8 uint = 65535
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            if v7 {
                v3[("" + (("hello" + ("hello" + "a")) + ("\n\t" + "")))] = 1
                if ((231 > 1) && (true && v7)) {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    v11 := GenS1{F0: "a"}
                    _ = v11
                }
            } else {
                var v12 float64 = (1.5 / v6)
                _ = v12
            }
            v13 := GenS2{F0: -1}
            fmt.Println(v7)
            _ = fi10
            _ = v13
        }
        _ = fi9
    }
    v14 := GenS3{F0: -1}
    if (0.1 != v6) {
        fmt.Println(v6, v1, v6)
        fmt.Println(v4)
    } else {
        switch 648 {
        case -241:
            v8 -= v8
        case 128:
            if (!(int16(v8) >= 387)) {
                delete(v3, ("foo bar" + ("" + "hello")))
                v15 := GenS1{F0: "\n\t"}
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            v4 -= (v1 >> ((146) & 63))
            for fi17 := 0; fi17 < 2; fi17++ {
                v8 = (v8 | v8)
                var v18 uint64 = 2
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint = v8
            v3["foo bar"] = 382
            _ = fi19
            _ = v20
        }
    }
    switch ("\n\t" + "") {
    case "":
        if false {
            fmt.Println(v8, v1)
            var v21 int8 = 1
            var v22 string = "a"
            _ = v21
            _ = v22
        } else {
            if (v7 && (!(v7 || (((747 != 0) || false) || false)))) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    fmt.Println(fi23, v7)
                    v8 = v8
                    v24 := GenS1{F0: "foo bar"}
                    _ = fi23
                    _ = v24
                }
                v3[(("" + "\n\t") + "a")] = 1
            } else {
                v6 -= v6
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 uint16 = 0
                    for fi27 := 0; fi27 < 3; fi27++ {
                        if true {
                            v28 := GenS4{F0: (fi25 << (((v8 % (((v8 % (((^(v8))) | 1))) | 1))) & 63)), F1: -9223372036854775808}
                            _ = v28
                        } else {
                            fmt.Println(v7, fi25, fi27)
                            v26 = v1
                            fmt.Println(v7, v4)
                        }
                        var v29 float64 = 3.14159
                        _ = fi27
                        _ = v29
                    }
                    _ = fi25
                    _ = v26
                }
                v30 := GenS4{F0: 42, F1: -9223372036854775808}
                _ = v30
            }
            v1 = v4
        }
    }
    if (!((32767 < 1) || (32767 <= 4294967295))) {
        fmt.Println(v8, v1)
    }
    var v31 float32 = 1e10
    for fi32 := 0; fi32 < 5; fi32++ {
        switch (("\n\t" + "hello") + "\n\t") {
        case "a":
            v33 := GenS3{F0: 2147483647}
            _ = v33
        case "\n\t":
            if (float64(3.14159) > (v6 + v6)) {
                v34 := GenS2{F0: int32(v4)}
                if v7 {
                    v35 := GenS3{F0: 2147483647}
                    _ = len(v2)
                    if (false && (true && (1 != 2))) {
                        v1 -= (^(v1))
                        var v36 float32 = (v31 - 0.0)
                        var v37 uint32 = 1
                        _ = v36
                        _ = v37
                    }
                    _ = v35
                } else {
                    var v38 uint32 = 4294967295
                    var v39 uint = v8
                    _ = v38
                    _ = v39
                }
                if v7 {
                    fmt.Println(v6, v1, v6)
                } else {
                    v4 -= v1
                }
                _ = v34
            } else {
                v4 -= v1
                v40 := v3[""]
                var v41 int32 = 0
                _ = v40
                _ = v41
            }
        case "foo bar":
            var v42 uint = (v8 >> ((v8) & 63))
            _ = v42
        default:
            fmt.Println(v6)
        }
        _ = fi32
    }
    switch 32767 {
    case 9223372036854775807:
        var v43 int32 = 204
        _ = v43
    }
    v7 = v7
    var v44 uint = (v8 - 1)
    v8 = v8
    v44 -= v44
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
