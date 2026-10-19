package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 bool }

func main() {
    var v1 string = ("a" + ("foo bar" + ("hello" + ((("hello" + ("a" + "a")) + "foo bar") + "\n\t"))))
    v2 := []int{32768, -9223372036854775808}
    var v3 int32 = -1
    var v4 int16 = 1
    var v5 int32 = 0
    var v6 int32 = (v5 + v5)
    switch 1 {
    case 42:
        for fi7 := 0; fi7 < 14; fi7++ {
            if len(v2) > 0 {
                v2[0] = fi7
            }
            _ = fi7
        }
    case 0:
        var v8 int16 = (v4 / ((v4) | 1))
        _ = v8
    case 915:
        if false {
            v6 = v6
            switch (v1 + "") {
            case "foo bar":
                var v9 float32 = (1.5 / 0.1)
                _ = v9
            }
            if true {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    var v10 string = v1
                    _ = v10
                }
            }
        } else {
            v11 := GenS1{F0: 0, F1: 42}
            var v12 uint64 = 1
            v13 := GenS1{F0: 539, F1: 4294967295}
            _ = v11
            _ = v12
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 10; fi14++ {
            v5 = (v3 / ((v3) | 1))
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 int64 = 81
                var v17 int64 = 42
                var v18 int8 = 42
                var v19 uint = 0
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi14
        }
    }
    v3 += (v6 ^ -2147483648)
    if ((2 == 32767) || (!(1 < 787))) {
        switch ("hello" + "\n\t") {
        case "hello":
            switch -9223372036854775808 {
            case 625:
                fmt.Println(v4)
            case 32767:
                if false {
                    v2 = append(v2, 529)
                }
            default:
                if (!true) {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        var v21 int32 = v6
                        _ = fi20
                        _ = v21
                    }
                    switch 4294967295 {
                    case 9223372036854775807:
                        v1 = ("\n\t" + "\n\t")
                    case 442:
                        var v22 int64 = 0
                        _ = v22
                    }
                    var v23 uint64 = 2
                    _ = v23
                }
            }
        case "foo bar":
            if false {
                var v24 uint16 = 255
                _ = v24
            } else {
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 string = (v1 + v1)
                    if ((v26 + "") <= "a") {
                        var v27 uint16 = 42
                        var v28 bool = false
                        _ = v27
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
                fmt.Println(v5)
            }
        default:
            var v29 uint = 897
            _ = v29
        }
        if true {
            switch v1 {
            case "\n\t":
                v30 := GenS1{F0: 1, F1: 2}
                _ = v30
            case "foo bar":
                for fi31 := 0; fi31 < 20; fi31++ {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        var v32 string = v1
                        _ = v32
                    case 42:
                        if false {
                            var v33 int16 = 549
                            var v34 int = 681
                            if true {
                                var v35 int = (fi31 | -1)
                                var v36 uint = 1
                                fmt.Println(v3, v4)
                                _ = v35
                                _ = v36
                            }
                            _ = v33
                            _ = v34
                        }
                    case -9223372036854775808:
                        v37 := GenS1{F0: 4294967295, F1: 400}
                        _ = v37
                    default:
                        v38 := GenS1{F0: 0, F1: 2147483647}
                        _ = v38
                    }
                    _ = fi31
                }
            case "hello":
                v3 = 0
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            switch v1 {
            case "":
                fmt.Println(v6, v5, v5)
            default:
                if false {
                    v1 = ""
                    if len(v2) > 0 {
                        v2[0] = -1
                    }
                    v4 = -1
                }
            }
        }
    }
    var v39 uint16 = 32767
    switch v1 {
    case "":
        var v40 int16 = v4
        _ = v40
    case "\n\t":
        for fi41 := 0; fi41 < 13; fi41++ {
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
            v6 = v3
            var v42 int8 = 127
            _ = fi41
            _ = v42
        }
    }
    v1 = v1
    v1 = "a"
    var v43 uint32 = 1
    var v44 string = v1
    v1 = ((v44 + v44) + v1)
    v1 = (v1 + "")
    var v45 uint16 = v39
    v3 = (v5 - 2147483647)
    if true {
        v46 := GenS1{F0: 947, F1: 4294967295}
        v39 = v45
        _ = v46
    } else {
        v47 := GenS2{F0: "hello", F1: true}
        _ = v47
    }
    switch "a" {
    case "hello":
        switch (v1 + v44) {
        case "":
            if (!true) {
                v6 = v5
                v48 := GenS1{F0: 656, F1: 1}
                _ = v48
            }
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v49 := GenS1{F0: 695, F1: v43}
            _ = v49
        }
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v2 = append(v2, 0)
    }
    var v50 int = 1
    v50 += 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v50
    _ = v6
}
