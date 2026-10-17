package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint64 }
func genFn0(p0 uint8, p1 uint16) uint64 {
    var v6 float32 = 3.14159
    var v7 uint64 = 1
    var v8 int = -767
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = ((1.5 / 1e10) >= 3.14159)
    var v3 int64 = 711
    var v4 bool = (!(true && (13 == 1)))
    v5 := make(map[string]string)
    v9 := GenS1{F0: 32767, F1: -1, F2: -2147483648}
    v10 := GenS1{F0: -2147483648, F1: -32768, F2: -915}
    var v11 uint8 = 127
    switch "a" {
    case "foo bar":
        switch 722 {
        case 4294967295:
            var v12 int64 = v3
            _ = v12
        default:
            v3 = v3
        }
    case "\n\t":
        fmt.Println(v11)
    case "a":
        delete(v1, 433)
    default:
        v13 := GenS1{F0: -126, F1: 458, F2: -2147483648}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            switch (^(fi14)) {
            case 658:
                var v16 float64 = (1e10 / float64(1.5))
                _ = v16
            case 124:
                fmt.Println(v2)
            case -39:
                v17 := GenS2{F0: 4294967295}
                _ = v17
            default:
                v18 := GenS1{F0: -932, F1: -32768, F2: 42}
                _ = v18
            }
            v4 = true
            delete(v5, "")
            _ = fi15
        }
        v3 = 399
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS2{F0: 0}
            var v21 uint64 = 4294967295
            v22 := []int{}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi14
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        switch (-(fi23)) {
        case 128:
            var v24 int16 = 32767
            _ = v24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        if (v2 && (!((100 == -125) && ((float32(0.1) <= 1.5) || (-32768 > -1))))) {
            v1[2147483648] = ("hello" + "\n\t")
        } else {
            v2 = (-4 >= -47)
            var v26 int16 = 834
            var v27 int8 = -1
            _ = v26
            _ = v27
        }
        v28 := GenS2{F0: 0}
        _ = fi25
        _ = v28
    }
    switch "foo bar" {
    case "hello":
        if (210 < 127) {
            for fi29 := 0; fi29 < 9; fi29++ {
                for fi30 := 0; fi30 < 1; fi30++ {
                    v31 := []int{fi30, 1}
                    var v32 uint8 = v11
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v33 := GenS2{F0: 0}
                if false {
                    v11 = v11
                    if (fi29 > fi29) {
                        if false {
                            var v34 uint16 = 1
                            var v35 int64 = (v3 >> ((949) & 63))
                            _ = v34
                            _ = v35
                        }
                        v36 := v1[184]
                        if (-69 != -1) {
                            switch (fi29 >> ((4294967295) & 63)) {
                            case 42:
                                var v37 uint32 = 65535
                                _ = v37
                            }
                        } else {
                            switch fi29 {
                            case 42:
                                var v38 int64 = v3
                                _ = v38
                            }
                        }
                        _ = v36
                    } else {
                        var v39 int32 = -226
                        _ = v39
                    }
                    v40 := []int{}
                    _ = v40
                }
                _ = fi29
                _ = v33
            }
            if (32768 == 4294967295) {
                var v41 int = 1
                for fi42 := 0; fi42 < 8; fi42++ {
                    v43 := GenS2{F0: 0}
                    for fi44 := 0; fi44 < 14; fi44++ {
                        fmt.Println(fi44)
                        var v45 int16 = -1
                        v11 = v11
                        _ = fi44
                        _ = v45
                    }
                    _ = fi42
                    _ = v43
                }
                for fi46 := 0; fi46 < 14; fi46++ {
                    for fi47 := 0; fi47 < 6; fi47++ {
                        v48 := GenS3{F0: 0, F1: v2}
                        v41 += fi46
                        _ = fi47
                        _ = v48
                    }
                    _ = fi46
                }
                _ = v41
            } else {
                for fi49 := 0; fi49 < 2; fi49++ {
                    fmt.Println(v2)
                    switch "\n\t" {
                    case "foo bar":
                        v2 = (42 != 2147483647)
                    }
                    _ = fi49
                }
                v4 = (-32768 < -397)
            }
        } else {
            if (-1 >= v3) {
                for fi50 := 0; fi50 < 4; fi50++ {
                    var v51 uint = 65535
                    var v52 float64 = 0.1
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                v53 := GenS4{F0: 979, F1: (^(v11))}
                delete(v5, "a")
                _ = v53
            }
            v54 := GenS5{F0: 1, F1: 895, F2: genFn0(128, 0)}
            if v4 {
                if v2 {
                    delete(v1, 1)
                    var v55 int32 = -2147483648
                    _ = len(v5)
                    _ = v55
                } else {
                    v56 := []int{-838, 9223372036854775807, -9223372036854775808}
                    var v57 int8 = 0
                    _ = v56
                    _ = v57
                }
                fmt.Println(v2, v3)
            }
            _ = v54
        }
    case "":
        v4 = (true && false)
    case "foo bar":
        v58 := GenS3{F0: 0, F1: (!v2)}
        _ = v58
    }
    v59 := v5[("\n\t" + (("hello" + ("foo bar" + ("a" + "a"))) + "foo bar"))]
    var v60 string = v59
    var v61 string = v60
    if true {
        var v62 uint16 = 65535
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v61
    _ = v9
}
