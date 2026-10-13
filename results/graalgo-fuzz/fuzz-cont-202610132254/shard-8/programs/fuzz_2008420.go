package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 bool, p2 uint) uint32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: "hello", F1: ("a" + ("foo bar" + "hello"))}
        for fi8 := 0; fi8 < 12; fi8++ {
            p1 = true
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    p1 = true
    _ = p0
    _ = p1
    _ = p2
    return 26
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 65535
    var v3 float32 = 1.5
    var v4 float64 = (3.14159 * -2.25)
    var v5 bool = (949 >= 65535)
    if false {
        if true {
            switch ("foo bar" + ("hello" + "")) {
            case "":
                if v5 {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v10 := []int{int(v1), fi9, (fi9 ^ fi9)}
                        _ = fi9
                        _ = v10
                    }
                    v5 = ((v4 * v4) < (0.0 * 1e10))
                } else {
                    var v11 uint8 = 164
                    v4 = v4
                    if v5 {
                        if v5 {
                            v12 := GenS1{F0: ("a" + "\n\t"), F1: ""}
                            _ = v12
                        }
                        v1 = (^(v1))
                        if (v5 || (true && (9223372036854775807 < 2147483647))) {
                            v13 := []int{32767, 0}
                            var v14 bool = (-9223372036854775808 == -532)
                            var v15 string = ("\n\t" + "foo bar")
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = v11
                }
            }
        } else {
            v2 += 0
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS2{F0: v2}
                v4 -= v4
                v3 += (v3 + v3)
                _ = fi16
                _ = v17
            }
            var v18 uint16 = (v2 - (^(v2)))
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            if true {
                var v20 int64 = -1
                var v21 string = ""
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                v23 := []int{(^(fi19)), (fi19 - fi22)}
                var v24 bool = v5
                var v25 uint32 = 173
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi19
        }
        var v26 uint32 = 268
        _ = v26
    }
    switch 0 {
    case 127:
        var v27 int32 = -2147483648
        _ = v27
    case 4294967295:
        v28 := GenS3{F0: 18}
        _ = v28
    case 42:
        if (2147483647 <= -642) {
            v29 := GenS1{F0: "", F1: (((("a" + "\n\t") + "hello") + "\n\t") + "\n\t")}
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := []int{1, fi30, 127}
                _ = fi30
                _ = v31
            }
            if (!true) {
                if false {
                    var v32 int16 = 0
                    _ = v32
                }
                for fi33 := 0; fi33 < 5; fi33++ {
                    v34 := GenS1{F0: "foo bar", F1: (("hello" + ("a" + "a")) + "")}
                    if v5 {
                        var v35 uint8 = 1
                        _ = v35
                    }
                    v36 := []int{fi33, fi33, (fi33 / ((fi33) | 1))}
                    _ = fi33
                    _ = v34
                    _ = v36
                }
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v4)
                }
            }
            _ = v29
        } else {
            for fi37 := 0; fi37 < 20; fi37++ {
                var v38 uint64 = 2
                var v39 int8 = 100
                _ = fi37
                _ = v38
                _ = v39
            }
            if true {
                var v40 uint64 = 0
                _ = v40
            }
        }
    }
    if false {
        v4 = v4
        v41 := GenS2{F0: (^(v2))}
        _ = v41
    } else {
        v2 = 37
        var v42 int64 = 1
        _ = v42
    }
    v4 = v4
    v2 = v2
    if (838 < 1) {
        if v5 {
            v43 := []int{2147483647}
            _ = v43
        }
    } else {
        v1 += v1
        var v44 bool = v5
        _ = v44
    }
    switch 225 {
    case -1:
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := []int{fi45, fi45, fi45}
            switch fi45 {
            case 32768:
                for fi47 := 0; fi47 < 6; fi47++ {
                    v48 := GenS3{F0: 2}
                    v49 := GenS1{F0: ("\n\t" + (("foo bar" + "") + "a")), F1: ""}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            }
            _ = fi45
            _ = v46
        }
    case 2147483647:
        if (v5 || v5) {
            var v50 float32 = -2.25
            _ = v50
        }
    case -719:
        for fi51 := 0; fi51 < 1; fi51++ {
            switch "a" {
            case "hello":
                fmt.Println(fi51)
            case "\n\t":
                var v52 uint = v1
                _ = v52
            case "foo bar":
                var v53 uint16 = 391
                _ = v53
            default:
                var v54 int16 = 1000
                _ = v54
            }
            v55 := []int{(-(fi51)), fi51}
            _ = fi51
            _ = v55
        }
    }
    v56 := GenS3{F0: 255}
    for fi57 := 0; fi57 < 7; fi57++ {
        var v58 uint32 = 2147483647
        var v59 float64 = 1e10
        _ = fi57
        _ = v58
        _ = v59
    }
    fmt.Println(v3)
    for fi60 := 0; fi60 < 14; fi60++ {
        for fi61 := 0; fi61 < 9; fi61++ {
            v62 := []int{42, fi60, (^(fi61))}
            var v63 uint = 42
            _ = fi61
            _ = v62
            _ = v63
        }
        v64 := GenS3{F0: 2}
        _ = fi60
        _ = v64
    }
    fmt.Println(v1, v3, v2)
    var v65 int = 4294967295
    v66 := GenS1{F0: "", F1: (("" + "hello") + ("hello" + ((("hello" + "foo bar") + "a") + ("\n\t" + ("" + "foo bar")))))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v65
    _ = v66
}
