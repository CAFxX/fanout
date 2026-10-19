package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    v4 := GenS1{F0: 255}
    if false {
        if true {
            var v5 uint = 4294967295
            var v6 float64 = 1.5
            _ = v5
            _ = v6
        }
        var v7 uint16 = 0
        v7 = 15
        _ = v7
    } else {
        v8 := []int{83}
        _ = v8
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    v1 := []string{("" + "\n\t")}
    var v2 string = ("a" + "")
    v3 := []uint32{42, 577, 42}
    fmt.Println(v2)
    fmt.Println(v2)
    switch 0 {
    case 32768:
        if len(v1) > 0 {
            v1[0] = (v2 + "")
        }
    case -781:
        if true {
            switch ("" + "a") {
            case "foo bar":
                if (1 < genFn0(1)) {
                    if (v2 <= (("" + (((v2 + v2) + v2) + v2)) + ("\n\t" + v2))) {
                        switch -9223372036854775808 {
                        case -2147483648:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 127:
                            v2 = v2
                        }
                    }
                    var v9 float64 = float64((3.14159 * 0.0))
                    if (2147483648 != 65535) {
                        switch 127 {
                        case 2147483647:
                            switch 127 {
                            case 32768:
                                var v10 uint32 = 0
                                _ = v10
                            case 238:
                                v11 := GenS2{F0: 127}
                                _ = v11
                            case 9223372036854775807:
                                var v12 int8 = 71
                                _ = v12
                            }
                        case -450:
                            v13 := GenS2{F0: 127}
                            _ = v13
                        case -40:
                            var v14 float32 = 0.0
                            _ = v14
                        default:
                            var v15 string = "a"
                            _ = v15
                        }
                    }
                    _ = v9
                }
            case "\n\t":
                var v16 uint = 44
                _ = v16
            case "a":
                for fi17 := 0; fi17 < 15; fi17++ {
                    v2 = v2
                    _ = fi17
                }
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 uint8 = 1
                v3 = append(v3, 65535)
                _ = fi18
                _ = v19
            }
        }
    case 42:
        var v20 int64 = 613
        _ = v20
    }
    v21 := GenS1{F0: 18446744073709551615}
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            switch (-(fi22)) {
            case 9223372036854775807:
                var v23 string = (("foo bar" + "\n\t") + v2)
                _ = v23
            case 1:
                var v24 uint64 = 18446744073709551615
                _ = v24
            case 127:
                for fi25 := 0; fi25 < 20; fi25++ {
                    v2 = v2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi25
                }
            }
            _ = len(v3)
        } else {
            var v26 string = v2
            _ = v26
        }
        _ = fi22
    }
    var v27 int = 875
    switch (v27 >> ((114) & 63)) {
    case 4294967295:
        v3 = append(v3, 196)
    case -1:
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(fi28)
            var v29 int16 = -1
            for fi30 := 0; fi30 < 4; fi30++ {
                if (("hello" == v2) || false) {
                    var v31 int8 = int8(v27)
                    if true {
                        var v32 uint32 = 1
                        switch v2 {
                        case "":
                            if (128 != 162) {
                                v33 := GenS1{F0: 196}
                                v29 -= (v29 / ((v29) | 1))
                                v2 = "hello"
                                v2 = v2
                                _ = v33
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = (v32 + v32)
                                }
                            }
                        case "\n\t":
                            v2 = "\n\t"
                        case "foo bar":
                            if (9223372036854775807 < 9223372036854775807) {
                                switch v27 {
                                case 440:
                                    v29 -= (-(v29))
                                    var v34 uint64 = 1
                                    var v35 uint32 = 675
                                    _ = v34
                                    _ = v35
                                case 4294967295:
                                    v32 = 1
                                }
                            }
                            var v36 bool = true
                            _ = v36
                        }
                        var v37 uint32 = 42
                        _ = v32
                        _ = v37
                    }
                    _ = v31
                } else {
                    v38 := GenS1{F0: 870}
                    v2 = "\n\t"
                    var v39 float32 = 0.0
                    _ = v38
                    _ = v39
                }
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
    case 32767:
        var v40 int64 = 988
        _ = v40
    default:
        v27 -= v27
    }
    v41 := GenS2{F0: 127}
    if false {
        v42 := GenS3{F0: (^(v27))}
        _ = v42
    } else {
        var v43 int16 = 1000
        var v44 string = v2
        var v45 int8 = -1
        _ = v43
        _ = v44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v41
}
