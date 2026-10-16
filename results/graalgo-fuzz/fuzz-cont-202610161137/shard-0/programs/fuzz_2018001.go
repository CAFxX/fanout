package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 int64 }
func genFn0() uint32 {
    var v4 float32 = float32(0.0)
    if (42 < 1) {
        var v5 string = ""
        _ = v5
    } else {
        v6 := []int{}
        _ = v6
    }
    _ = v4
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]uint32)
    var v3 uint16 = 0
    for fi7 := 0; fi7 < 14; fi7++ {
        v1 -= (v1 | 2147483648)
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    switch int(v3) {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v1, fi9, v1)
            _ = fi9
        }
    case 32767:
        var v10 uint64 = 4294967295
        _ = v10
    case 127:
        v11 := []int{}
        _ = v11
    }
    v12 := GenS1{F0: "", F1: true, F2: 4294967295}
    switch -2147483648 {
    case 0:
        v3 = (^(v3))
    case -1:
        var v13 int = 890
        _ = v13
    case 32768:
        v14 := []int{128, 128, 127, 42}
        _ = v14
    }
    v3 = (^(v3))
    switch 2147483647 {
    case 42:
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 int8 = -128
                _ = fi16
                _ = v17
            }
            v18 := []int{fi15, (fi15 ^ 32767), fi15, 127}
            _ = fi15
            _ = v18
        }
    case -1:
        v19 := GenS1{F0: "a", F1: true, F2: 1}
        _ = v19
    case 32767:
        switch 603 {
        case 32767:
            v20 := GenS1{F0: "", F1: false, F2: uint(v3)}
            _ = v20
        }
    default:
        if false {
            fmt.Println(v1, v1)
            var v21 bool = (false || (!false))
            for fi22 := 0; fi22 < 3; fi22++ {
                v3 = (v3 / ((159) | 1))
                v23 := GenS1{F0: "", F1: v21, F2: 149}
                _ = fi22
                _ = v23
            }
            _ = v21
        }
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := v2["\n\t"]
        if (9223372036854775807 != -9223372036854775808) {
            fmt.Println(v25, v25)
        }
        v26 := GenS1{F0: ("" + "a"), F1: true, F2: 716}
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            if false {
                fmt.Println(fi28, fi28, fi28)
                var v29 int = fi28
                if false {
                    var v30 uint64 = 997
                    delete(v2, "foo bar")
                    var v31 float64 = ((1e10 - (-2.25 * float64(1.5))) / (3.14159 * 0.1))
                    fmt.Println(v30, v1)
                    _ = v30
                    _ = v31
                }
                _ = v29
            }
            if true {
                switch "" {
                case "":
                    var v32 uint = 255
                    _ = v32
                }
                if false {
                    v1 = (v1 << ((875) & 63))
                    switch fi27 {
                    case -717:
                        v3 = (v3 - 65535)
                    case -2147483648:
                        v1 = (^(v1))
                    case 115:
                        v33 := GenS1{F0: "a", F1: true, F2: 42}
                        _ = v33
                    default:
                        fmt.Println(fi28, fi27)
                    }
                }
            }
            var v34 uint8 = 127
            if false {
                v3 = v3
                v35 := GenS1{F0: "", F1: ("" != "\n\t"), F2: 255}
                _ = v35
            } else {
                var v36 uint = uint(v3)
                v34 = (^(v34))
                _ = v36
            }
            v37 := GenS1{F0: "", F1: (1000 > -1), F2: 1}
            fmt.Println(fi27, fi28, fi27)
            _ = fi28
            _ = v34
            _ = v37
        }
        for fi38 := 0; fi38 < 3; fi38++ {
            switch "\n\t" {
            case "":
                v39 := GenS1{F0: "a", F1: false, F2: 187}
                _ = v39
            default:
                if ((1000 <= 858) || (!true)) {
                    v40 := []int{fi27}
                    if len(v40) > 0 {
                        v40[0] = fi38
                    }
                    _ = v40
                }
            }
            var v41 uint16 = (v3 % (((v3 / (((^(v3))) | 1))) | 1))
            fmt.Println(v3)
            _ = fi38
            _ = v41
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            v43 := GenS1{F0: ((("\n\t" + "hello") + ("a" + "\n\t")) + ("" + "")), F1: true, F2: 1}
            v44 := GenS1{F0: "a", F1: (1.5 < -2.25), F2: 1}
            v45 := []int{(fi27 >> ((208) & 63))}
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        _ = fi27
    }
    if (!(1e10 != 1e10)) {
        v46 := GenS1{F0: ((("" + "a") + ("\n\t" + ("a" + ("foo bar" + ("\n\t" + ("hello" + "hello")))))) + ("a" + (((("" + (("" + "\n\t") + "a")) + ("foo bar" + (("\n\t" + ((("hello" + ("hello" + "hello")) + "a") + ("" + ""))) + ("a" + "hello")))) + "foo bar") + "\n\t"))), F1: false, F2: 1}
        for fi47 := 0; fi47 < 11; fi47++ {
            v48 := GenS2{F0: 1}
            var v49 int16 = 42
            if false {
                for fi50 := 0; fi50 < 10; fi50++ {
                    fmt.Println(v1, v3)
                    _ = fi50
                }
                if true {
                    fmt.Println(v49)
                    for fi51 := 0; fi51 < 14; fi51++ {
                        v3 += 2
                        switch "\n\t" {
                        case "hello":
                            var v52 int64 = -664
                            _ = v52
                        case "":
                            v53 := GenS3{F0: 0, F1: -32768, F2: -1}
                            _ = v53
                        }
                        fmt.Println(v1, fi47, fi47)
                        _ = fi51
                    }
                    var v54 int32 = 1
                    _ = v54
                }
                if false {
                    var v55 string = "\n\t"
                    v56 := []int{fi47, (fi47 - fi47)}
                    _ = v55
                    _ = v56
                } else {
                    switch (fi47 | fi47) {
                    case 42:
                        for fi57 := 0; fi57 < 3; fi57++ {
                            v49 += (v49 | (v49 & 1000))
                            _ = fi57
                        }
                    case 127:
                        fmt.Println(v3, fi47, v1)
                    case 27:
                        fmt.Println(v3)
                    }
                }
            }
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v46
    }
    switch 298 {
    case 2147483647:
        v58 := GenS3{F0: -1, F1: 1, F2: -9223372036854775808}
        _ = v58
    }
    v3 -= (v3 * (v3 & 1))
    if false {
        if false {
            v59 := []int{945, -114, 127, -1}
            _ = v59
        }
    }
    v2["foo bar"] = (^(v1))
    var v60 int16 = -32768
    var v61 uint8 = 156
    if ((v3 % (((^(v3))) | 1)) != (v3 << ((42) & 63))) {
        v62 := []int{181, 32767}
        _ = v62
    }
    var v63 int = 9223372036854775807
    for fi64 := 0; fi64 < 2; fi64++ {
        switch "a" {
        case "":
            v65 := GenS2{F0: 705}
            _ = v65
        }
        fmt.Println(v1, v61, fi64)
        var v66 string = "\n\t"
        _ = fi64
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v63
}
