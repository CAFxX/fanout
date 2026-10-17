package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 bool) int8 {
    var v7 int64 = 354
    var v8 int16 = -1
    var v9 bool = false
    v9 = false
    switch 127 {
    case 32768:
        var v10 int8 = 0
        _ = v10
    case 162:
        v11 := []int{-9223372036854775808}
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -80
}
func genFn1() bool {
    v12 := []int{2147483647, -583, -919, 0}
    if true {
        var v13 uint64 = 18446744073709551615
        var v14 uint = 42
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            _ = len(v12)
            _ = fi15
        }
        if false {
            var v16 uint = 0
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 string = "a"
                _ = fi17
                _ = v18
            }
            if (true && ((!true) && true)) {
                v16 = (v16 / ((v16) | 1))
                if len(v12) > 0 {
                    v12[0] = -1
                }
            }
            _ = v16
        }
    }
    if true {
        var v19 bool = (true || false)
        _ = v19
    } else {
        if (!false) {
            if len(v12) > 0 {
                v12[0] = 127
            }
            if (true || true) {
                if true {
                    var v20 uint32 = 4294967295
                    switch "" {
                    case "hello":
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case "a":
                        v20 -= 0
                    case "\n\t":
                        if (("a" < (("\n\t" + ((("foo bar" + "foo bar") + "") + "a")) + "")) && true) {
                            for fi21 := 0; fi21 < 4; fi21++ {
                                v20 = 65535
                                _ = fi21
                            }
                            for fi22 := 0; fi22 < 17; fi22++ {
                                fmt.Println(v20, fi22)
                                fmt.Println(v20)
                                var v23 uint32 = 1
                                _ = fi22
                                _ = v23
                            }
                            var v24 uint8 = 42
                            _ = v24
                        }
                    }
                    var v25 int = 42
                    _ = v20
                    _ = v25
                } else {
                    var v26 float64 = 0.1
                    v26 = v26
                    fmt.Println(v26)
                    _ = v26
                }
            }
        }
    }
    _ = v12
    return (18446744073709551615 > 0)
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 2
    v3 := make(map[int]int)
    var v4 bool = (4294967295 < 42)
    var v5 string = "foo bar"
    v6 := make(map[int]int)
    var v27 float32 = 3.14159
    if (42 != 2147483647) {
        for fi28 := 0; fi28 < 7; fi28++ {
            v4 = (false || (-318 <= 2147483647))
            if ("\n\t" > (((v5 + v5) + (v5 + (v5 + "hello"))) + (v5 + (v5 + "foo bar")))) {
                v1 = (v1 * (v1 * v1))
                if v4 {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v4 = v4
                        v30 := []int{127, (fi28 * fi29)}
                        switch "hello" {
                        case "foo bar":
                            v31 := GenS1{F0: 0, F1: 1, F2: (((v5 + v5) + ((v5 + ("hello" + (v5 + ("hello" + v5)))) + (v5 + v5))) + ("hello" + v5))}
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                } else {
                    fmt.Println(v4, v1)
                }
            } else {
                for fi32 := 0; fi32 < 16; fi32++ {
                    if (v4 || v4) {
                        var v33 int16 = 32767
                        v2 = v2
                        switch (fi28 * (-(fi32))) {
                        case -2147483648:
                            switch fi28 {
                            case -554:
                                switch (fi28 << ((1) & 63)) {
                                case 128:
                                    v27 -= v27
                                    v5 = v5
                                    var v34 float32 = v27
                                    _ = v34
                                }
                            case 4294967295:
                                fmt.Println(v5)
                            case 9223372036854775807:
                                fmt.Println(fi32, v4, v5)
                            default:
                                v33 = -1
                            }
                        case 42:
                            var v35 uint8 = 42
                            _ = v35
                        case 1:
                            v36 := v3[(^(fi28))]
                            _ = v36
                        default:
                            var v37 uint8 = 19
                            _ = v37
                        }
                        v5 = v5
                        v38 := GenS2{F0: -52, F1: (-(v1))}
                        _ = v33
                        _ = v38
                    } else {
                        v5 = v5
                        var v39 uint8 = 127
                        _ = v39
                    }
                    v40 := GenS1{F0: 42, F1: 491, F2: "\n\t"}
                    _ = fi32
                    _ = v40
                }
            }
            _ = fi28
        }
    } else {
        switch -442 {
        case 361:
            var v41 uint64 = 54
            _ = v41
        case 0:
            v1 = v1
        default:
            v42 := GenS2{F0: 0, F1: v1}
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 20; fi43++ {
        var v44 int32 = 42
        v45 := GenS3{F0: (v2 >> ((v2) & 63)), F1: 42, F2: v1}
        if (fi43 < int(v1)) {
            var v46 int8 = 127
            _ = v46
        }
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        if v4 {
            fmt.Println(v2)
        }
        v48 := GenS2{F0: genFn0((-32768 < 0)), F1: v1}
        _ = fi47
        _ = v48
    }
    var v49 float64 = float64(v27)
    var v50 uint8 = 1
    for fi51 := 0; fi51 < 4; fi51++ {
        for fi52 := 0; fi52 < 6; fi52++ {
            var v53 string = v5
            _ = fi52
            _ = v53
        }
        _ = fi51
    }
    v27 = ((1.5 / v27) / ((1.5 / (-2.25 - v27)) / 1.5))
    v2 -= v2
    fmt.Println(v2, v1)
    for fi54 := 0; fi54 < 14; fi54++ {
        var v55 float32 = v27
        var v56 int32 = -658
        var v57 uint8 = v50
        _ = fi54
        _ = v55
        _ = v56
        _ = v57
    }
    var v58 int = 482
    var v59 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v58
    _ = v59
    _ = v6
}
