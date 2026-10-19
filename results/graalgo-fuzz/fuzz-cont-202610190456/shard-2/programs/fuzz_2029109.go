package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int16 {
    v6 := GenS1{F0: -474}
    if true {
        switch (("a" + "foo bar") + "a") {
        case "a":
            var v7 string = "a"
            _ = v7
        case "\n\t":
            var v8 int16 = 32767
            _ = v8
        case "foo bar":
            v9 := []int{9223372036854775807}
            _ = v9
        }
        switch 1 {
        case 128:
            switch 4294967295 {
            case -2147483648:
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint16 = 65535
                    for fi12 := 0; fi12 < 6; fi12++ {
                        fmt.Println(v11, v11)
                        _ = fi12
                    }
                    var v13 int16 = -600
                    _ = fi10
                    _ = v11
                    _ = v13
                }
            case 42:
                var v14 uint64 = 42
                _ = v14
            default:
                var v15 uint = 847
                _ = v15
            }
        case -9223372036854775808:
            if false {
                var v16 uint8 = 1
                if false {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        switch (-(fi17)) {
                        case 128:
                            if (354 <= 42) {
                                v16 = (v16 * (v16 / (((v16 & 42)) | 1)))
                            } else {
                                var v18 float32 = 1e10
                                v16 -= v16
                                _ = v18
                            }
                        case 42:
                            if true {
                                var v19 uint16 = 42
                                v19 = (v19 + v19)
                                _ = v19
                            }
                        case -9223372036854775808:
                            var v20 uint8 = 255
                            _ = v20
                        }
                        var v21 uint8 = v16
                        var v22 float32 = 3.14159
                        _ = fi17
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 18; fi23++ {
                        fmt.Println(v16)
                        var v24 float32 = (3.14159 / (0.1 - float32(3.14159)))
                        v16 = 0
                        _ = fi23
                        _ = v24
                    }
                }
                v25 := GenS1{F0: 9223372036854775807}
                _ = v16
                _ = v25
            } else {
                var v26 int = 0
                switch (((("foo bar" + "\n\t") + "hello") + "") + ((("\n\t" + "\n\t") + "foo bar") + "\n\t")) {
                case "foo bar":
                    v26 = (v26 | v26)
                case "a":
                    var v27 int64 = 0
                    _ = v27
                }
                v26 = v26
                _ = v26
            }
        default:
            switch "hello" {
            case "foo bar":
                var v28 uint8 = 46
                _ = v28
            case "hello":
                var v29 int16 = -32768
                _ = v29
            }
        }
    }
    _ = v6
    return -32768
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    v4 := []string{("hello" + "hello")}
    v5 := []int{127, -358, 111}
    v4 = append(v4, ("foo bar" + ("" + "foo bar")))
    var v30 uint8 = 255
    v3[(("hello" + ("hello" + "a")) + ("" + "foo bar"))] = 4294967295
    v30 -= (v30 & v30)
    switch ("hello" + "") {
    case "hello":
        for fi31 := 0; fi31 < 14; fi31++ {
            fmt.Println(v30)
            v1 += (v1 >> ((2) & 63))
            v1 = v1
            _ = fi31
        }
    case "\n\t":
        for fi32 := 0; fi32 < 19; fi32++ {
            _ = len(v5)
            _ = fi32
        }
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        for fi34 := 0; fi34 < 17; fi34++ {
            v1 -= (v1 % (((v1 ^ uint32(v30))) | 1))
            var v35 float64 = (-2.25 + ((1e10 - (1e10 / 0.1)) / 0.1))
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    if ((false || true) || true) {
        fmt.Println(v1)
    }
    switch -1 {
    case 2147483647:
        if true {
            switch ("foo bar" + ("" + "foo bar")) {
            case "":
                v30 -= (v30 & (v30 - (v30 << ((41) & 63))))
            case "a":
                var v36 float64 = 1.5
                _ = v36
            default:
                v37 := GenS1{F0: 32768}
                _ = v37
            }
            var v38 uint16 = 0
            v39 := GenS1{F0: 9223372036854775807}
            _ = v38
            _ = v39
        }
    case 127:
        var v40 uint8 = v30
        _ = v40
    case 32767:
        for fi41 := 0; fi41 < 5; fi41++ {
            for fi42 := 0; fi42 < 10; fi42++ {
                v30 += 128
                switch ("" + "a") {
                case "hello":
                    var v43 uint64 = 2
                    _ = v43
                default:
                    if false {
                        v30 = v30
                        var v44 int16 = genFn0()
                        _ = v44
                    } else {
                        var v45 int64 = -438
                        _ = v45
                    }
                }
                _ = fi42
            }
            var v46 uint = 2
            var v47 string = ""
            _ = fi41
            _ = v46
            _ = v47
        }
    default:
        fmt.Println(v30)
    }
    v48 := GenS1{F0: 128}
    v49 := GenS1{F0: -159}
    v50 := GenS1{F0: 237}
    v51 := GenS1{F0: -714}
    v52 := GenS1{F0: -1}
    var v53 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
