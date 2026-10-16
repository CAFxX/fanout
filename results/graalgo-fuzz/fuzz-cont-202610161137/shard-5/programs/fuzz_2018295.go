package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int16 }
func genFn0() int16 {
    var v5 uint8 = 0
    if false {
        var v6 uint64 = 18446744073709551615
        _ = v6
    } else {
        v7 := GenS1{F0: 2147483647, F1: 490}
        fmt.Println(v5)
        _ = v7
    }
    _ = v5
    return -719
}

func main() {
    v1 := []int{128, 32767}
    var v2 bool = (255 >= 65535)
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v8 int8 = 127
    v9 := GenS1{F0: 1, F1: 4294967295}
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint64 = 18446744073709551615
        switch (fi10 ^ (-(fi10))) {
        case 60:
            fmt.Println(v2)
        case -358:
            var v12 string = (("a" + "foo bar") + ("a" + (((("" + ("" + "\n\t")) + "hello") + "\n\t") + "hello")))
            _ = v12
        case 4294967295:
            for fi13 := 0; fi13 < 12; fi13++ {
                v2 = (!v2)
                _ = fi13
            }
        }
        _ = fi10
        _ = v11
    }
    v8 += -1
    for fi14 := 0; fi14 < 18; fi14++ {
        delete(v3, "a")
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
        v16 := v3["foo bar"]
        _ = fi14
        _ = v16
    }
    v8 = v8
    switch 2147483647 {
    case 128:
        v17 := GenS2{F0: 1000}
        _ = v17
    case 4294967295:
        v18 := GenS3{F0: 221}
        _ = v18
    case 127:
        var v19 uint64 = 666
        _ = v19
    }
    var v20 uint16 = 65535
    for fi21 := 0; fi21 < 2; fi21++ {
        if (!(!true)) {
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS2{F0: -32768}
                var v24 float64 = -2.25
                v25 := GenS4{F0: v2, F1: 1}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v26 uint = 1
            _ = v26
        } else {
            fmt.Println(v2)
            switch (fi21 * (fi21 ^ 32767)) {
            case 346:
                var v27 uint8 = 127
                _ = v27
            case -9223372036854775808:
                _ = len(v1)
            }
            for fi28 := 0; fi28 < 5; fi28++ {
                v2 = v2
                _ = fi28
            }
        }
        v29 := GenS5{F0: v2, F1: 1}
        if (v2 || ((^(v20)) == v20)) {
            fmt.Println(fi21, v8, v2)
            if (uint8(v20) < 249) {
                switch 32768 {
                case 0:
                    v8 -= v8
                case 32767:
                    var v30 uint64 = 2
                    _ = v30
                case 128:
                    for fi31 := 0; fi31 < 14; fi31++ {
                        v32 := GenS5{F0: v2, F1: 1}
                        fmt.Println(fi31, v20, v2)
                        var v33 int64 = 384
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                }
                if (!(!false)) {
                    var v34 uint16 = (v20 >> ((888) & 63))
                    for fi35 := 0; fi35 < 10; fi35++ {
                        _ = len(v1)
                        v8 += v8
                        _ = fi35
                    }
                    _ = v34
                }
            } else {
                for fi36 := 0; fi36 < 8; fi36++ {
                    var v37 float64 = 0.0
                    v8 -= v8
                    v38 := GenS1{F0: 4294967295, F1: (fi36 >> ((763) & 63))}
                    if true {
                        delete(v3, "\n\t")
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            var v39 int8 = int8(v20)
                            _ = v39
                        case "a":
                            v40 := GenS4{F0: (fi21 != 0), F1: 1000}
                            _ = v40
                        case "\n\t":
                            fmt.Println(v20)
                        default:
                            var v41 bool = (1 != 42)
                            _ = v41
                        }
                        fmt.Println(v20, v8, v37)
                        var v42 int32 = 215
                        _ = v42
                    } else {
                        switch (fi21 - -9223372036854775808) {
                        case -2147483648:
                            v43 := GenS6{F0: (fi21 ^ -837), F1: 128, F2: 1}
                            _ = v43
                        case 42:
                            var v44 uint32 = 1
                            _ = v44
                        case 1:
                            if v2 {
                                fmt.Println(v8, fi36, v8)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                switch "foo bar" {
                                case "":
                                    fmt.Println(fi36, v37)
                                case "a":
                                    var v45 int = 501
                                    _ = v45
                                default:
                                    v2 = false
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v46 uint16 = (^(v20))
                                _ = v46
                            }
                        default:
                            var v47 uint = 0
                            _ = v47
                        }
                    }
                    if false {
                        var v48 uint8 = 128
                        var v49 int8 = (^(v8))
                        var v50 int8 = (v49 ^ v8)
                        _ = v48
                        _ = v49
                        _ = v50
                    } else {
                        switch (("a" + "\n\t") + "") {
                        case "a":
                            v51 := GenS3{F0: (v20 ^ v20)}
                            var v52 bool = v2
                            if ((v20 % (((v20 & v20)) | 1)) != 655) {
                                var v53 int64 = 0
                                var v54 bool = v52
                                fmt.Println(v20)
                                _ = v53
                                _ = v54
                            }
                            _ = v51
                            _ = v52
                        case "foo bar":
                            var v55 bool = v2
                            _ = v55
                        }
                        v8 = (v8 >> ((42) & 63))
                    }
                    var v56 uint16 = (v20 ^ v20)
                    var v57 int8 = (v8 & (v8 >> ((65535) & 63)))
                    v58 := GenS7{F0: 42}
                    v8 -= v8
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v56
                    _ = v57
                    _ = v58
                }
            }
        } else {
            fmt.Println(v20, fi21, v2)
            var v59 int = (fi21 - (fi21 ^ 250))
            v2 = (42 == 42)
            _ = v59
        }
        _ = fi21
        _ = v29
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        var v61 float32 = 3.14159
        _ = fi60
        _ = v61
    }
    var v62 uint8 = 2
    fmt.Println(v20, v8, v62)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v62
    _ = v8
    _ = v9
}
