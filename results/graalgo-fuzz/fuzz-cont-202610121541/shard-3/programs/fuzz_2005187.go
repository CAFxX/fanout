package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint16 }
func genFn0() int {
    var v7 uint32 = 65535
    if true {
        var v8 uint16 = 32767
        _ = v8
    } else {
        v9 := GenS1{F0: 1}
        switch 32768 {
        case -1:
            v10 := GenS1{F0: 42}
            _ = v10
        case -971:
            var v11 string = ("" + "")
            _ = v11
        case 0:
            fmt.Println(v7)
        default:
            fmt.Println(v7)
        }
        v7 -= (v7 % ((v7) | 1))
        _ = v9
    }
    var v12 bool = (1e10 <= (-2.25 + -2.25))
    if (9223372036854775807 < 0) {
        var v13 string = "hello"
        fmt.Println(v12)
        _ = v13
    } else {
        var v14 int = 616
        var v15 int64 = -294
        v12 = true
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int = (fi16 << ((uint(v7)) & 63))
        switch (("a" + "\n\t") + ((("foo bar" + "\n\t") + "") + ((("\n\t" + "foo bar") + "hello") + ""))) {
        case "":
            var v18 uint16 = 511
            _ = v18
        case "\n\t":
            switch 2147483647 {
            case 859:
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch 2147483647 {
                    case 32767:
                        switch "\n\t" {
                        case "foo bar":
                            var v20 string = "\n\t"
                            _ = v20
                        case "":
                            var v21 float64 = 1.5
                            _ = v21
                        }
                    default:
                        switch 4294967295 {
                        case 2147483647:
                            v22 := []int{fi16, 1, 2147483647, v17}
                            _ = v22
                        case 0:
                            v23 := GenS1{F0: 65535}
                            _ = v23
                        case 127:
                            v12 = (!(v12 && (-859 >= 0)))
                        default:
                            v24 := GenS2{F0: "hello"}
                            _ = v24
                        }
                    }
                    if false {
                        v7 -= (v7 - 2)
                        var v25 uint16 = 255
                        fmt.Println(v7, v25)
                        _ = v25
                    } else {
                        v26 := []int{v17, (fi19 - fi19), (-(fi19))}
                        _ = v26
                    }
                    _ = fi19
                }
            default:
                for fi27 := 0; fi27 < 5; fi27++ {
                    switch fi27 {
                    case 1:
                        var v28 int64 = -1
                        _ = v28
                    case -2147483648:
                        v29 := []int{(fi16 - fi27)}
                        _ = v29
                    case -291:
                        fmt.Println(fi27, fi16, v17)
                    }
                    _ = fi27
                }
            }
        }
        v30 := GenS3{F0: -2147483648}
        _ = fi16
        _ = v17
        _ = v30
    }
    _ = v12
    _ = v7
    return 232
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]int)
    var v3 float32 = -2.25
    v4 := []uint32{0, 88, 1}
    var v5 string = (("hello" + ("\n\t" + "foo bar")) + (("hello" + "\n\t") + "\n\t"))
    var v6 int32 = -2147483648
    v31 := v2[(("a" + ("a" + v5)) + v5)]
    v3 = ((v3 - v3) + 1e10)
    if false {
        switch (("a" + (v5 + v5)) + v5) {
        case "foo bar":
            _ = len(v4)
        case "a":
            if true {
                var v32 int16 = 1
                if false {
                    delete(v2, "foo bar")
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    var v33 bool = false
                    _ = v33
                }
                var v34 uint64 = (v1 ^ 655)
                _ = v32
                _ = v34
            }
        }
    }
    v35 := GenS4{F0: 803, F1: -1}
    v36 := GenS5{F0: 386, F1: (v5 + v5)}
    if (-187 == -759) {
        for fi37 := 0; fi37 < 8; fi37++ {
            if false {
                if true {
                    fmt.Println(fi37)
                }
            } else {
                v38 := GenS4{F0: 42, F1: -104}
                if false {
                    v6 = v6
                    if (-1 <= v6) {
                        var v39 int64 = -9223372036854775808
                        _ = v39
                    } else {
                        v5 = v5
                    }
                    var v40 int16 = 1000
                    _ = v40
                } else {
                    var v41 int8 = 1
                    var v42 float64 = 1e10
                    if (65535 < 0) {
                        for fi43 := 0; fi43 < 9; fi43++ {
                            v41 = v41
                            _ = fi43
                        }
                    } else {
                        v6 = int32(v41)
                    }
                    _ = v41
                    _ = v42
                }
                _ = v38
            }
            v1 -= (v1 + 1)
            _ = fi37
        }
    } else {
        var v44 float64 = (-2.25 - -2.25)
        var v45 float64 = 1e10
        var v46 uint64 = (v1 >> ((42) & 63))
        _ = v44
        _ = v45
        _ = v46
    }
    var v47 float64 = 0.0
    _ = len(v2)
    var v48 int32 = (v6 ^ (v6 - v6))
    v4 = append(v4, 664)
    switch "foo bar" {
    case "\n\t":
        if true {
            v47 -= (v47 - v47)
            if (v1 <= (^(v1))) {
                v49 := GenS6{F0: 32}
                _ = v49
            }
        }
    case "":
        v5 = ("hello" + ((((((v5 + "") + v5) + ("foo bar" + "a")) + "\n\t") + ("hello" + (v5 + v5))) + v5))
    case "a":
        for fi50 := 0; fi50 < 5; fi50++ {
            switch "foo bar" {
            case "hello":
                v51 := GenS4{F0: 42, F1: 44}
                _ = v51
            case "foo bar":
                for fi52 := 0; fi52 < 10; fi52++ {
                    var v53 int32 = (v48 & 1)
                    fmt.Println(fi52, v53, v47)
                    _ = fi52
                    _ = v53
                }
            }
            switch fi50 {
            case -1:
                var v54 bool = false
                _ = v54
            default:
                _ = len(v4)
            }
            _ = fi50
        }
    default:
        v55 := GenS2{F0: "hello"}
        _ = v55
    }
    v31 -= (v31 ^ 4294967295)
    var v56 uint16 = 255
    var v57 int8 = -62
    if false {
        var v58 uint64 = 936
        for fi59 := 0; fi59 < 8; fi59++ {
            v6 += v48
            _ = fi59
        }
        _ = v58
    }
    if ((v6 >> ((255) & 63)) == v6) {
        switch v31 {
        case -1:
            v60 := GenS1{F0: 2}
            _ = v60
        case 32768:
            var v61 int64 = 91
            _ = v61
        }
        var v62 int32 = 0
        _ = v62
    }
    if (v31 < (v31 & genFn0())) {
        var v63 uint = 255
        v64 := GenS3{F0: 0}
        v6 += (-(v6))
        _ = v63
        _ = v64
    }
    v1 -= 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
