package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint64; F2 int }
type GenS6 struct { F0 bool; F1 string; F2 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 uint; F2 int64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64 }
func genFn0(p0 int32, p1 string) int {
    for fi5 := 0; fi5 < 8; fi5++ {
        if false {
            v6 := GenS1{F0: 365, F1: 1, F2: 42}
            for fi7 := 0; fi7 < 12; fi7++ {
                p0 = -576
                _ = fi7
            }
            _ = v6
        } else {
            v8 := GenS2{F0: 208}
            _ = v8
        }
        var v9 int16 = 32767
        for fi10 := 0; fi10 < 12; fi10++ {
            if (696 <= 931) {
                v11 := GenS3{F0: p0, F1: (-2.25 <= 0.0)}
                var v12 int32 = p0
                _ = v11
                _ = v12
            }
            v9 = (v9 * 1)
            _ = fi10
        }
        _ = fi5
        _ = v9
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            p1 = ("\n\t" + "hello")
            p0 -= (p0 * p0)
            p0 = -1
            _ = fi14
        }
        p1 = ("\n\t" + "a")
        var v15 int16 = -572
        _ = fi13
        _ = v15
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    var v3 int8 = 127
    var v4 int8 = (v3 + v3)
    if (!(65535 <= 487)) {
        v4 += (^(v4))
        if (!false) {
            if false {
                v4 = (v3 ^ v4)
                v2["hello"] = 0
                v16 := GenS2{F0: 42}
                _ = v16
            }
            v17 := GenS4{F0: ("\n\t" + "foo bar")}
            fmt.Println(v1)
            _ = v17
        } else {
            v18 := GenS2{F0: 584}
            _ = v18
        }
        if ((true && (194 >= 41)) && (!(!true))) {
            var v19 uint8 = 255
            _ = v19
        } else {
            if (174 >= -561) {
                if true {
                    var v20 string = "a"
                    v21 := []int{454, 42}
                    _ = v20
                    _ = v21
                }
                v1 = (v1 ^ (-(v1)))
                var v22 float32 = (0.1 / 0.1)
                _ = v22
            } else {
                for fi23 := 0; fi23 < 17; fi23++ {
                    switch "\n\t" {
                    case "foo bar":
                        v24 := GenS5{F0: 4294967295, F1: 1, F2: (fi23 - (fi23 * (fi23 + fi23)))}
                        _ = v24
                    case "a":
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v3 = v4
                            fmt.Println(v1)
                            var v26 uint = 2
                            v27 := GenS6{F0: false, F1: "hello", F2: 2}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    case "hello":
                        v28 := []int{fi23}
                        _ = v28
                    }
                    if (286 <= 32767) {
                        v29 := GenS4{F0: "hello"}
                        var v30 uint32 = 2
                        _ = v29
                        _ = v30
                    }
                    _ = fi23
                }
                for fi31 := 0; fi31 < 6; fi31++ {
                    v32 := GenS7{F0: 0}
                    if false {
                        if true {
                            switch -2147483648 {
                            case 128:
                                var v33 int8 = (v4 * -1)
                                _ = v33
                            case 4294967295:
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    switch (fi31 << ((0) & 63)) {
                                    case 1:
                                        var v35 uint16 = 883
                                        _ = v35
                                    case 0:
                                        switch "\n\t" {
                                        case "":
                                            if false {
                                                var v36 int = 32767
                                                _ = v36
                                            } else {
                                                v1 = (v1 % (((v1 & v1)) | 1))
                                            }
                                            v37 := []int{fi31}
                                            _ = v37
                                        case "hello":
                                            fmt.Println(fi31, fi31)
                                        case "foo bar":
                                            switch fi34 {
                                            case 9223372036854775807:
                                                var v38 uint8 = 128
                                                _ = v38
                                            }
                                        }
                                    }
                                    _ = fi34
                                }
                            }
                            v1 = v1
                            switch -2147483648 {
                            case 2147483647:
                                v39 := GenS8{F0: (true && (!((1e10 + ((0.1 + 0.0) * 1.5)) >= float64((float32(-2.25) + 1.5))))), F1: 255, F2: -849}
                                _ = v39
                            case 32768:
                                v2[("\n\t" + ("a" + "hello"))] = 225
                            }
                        }
                    }
                    _ = fi31
                    _ = v32
                }
            }
            v1 = 42
            switch ("\n\t" + (("\n\t" + "\n\t") + "foo bar")) {
            case "a":
                v40 := []int{-2147483648}
                _ = v40
            case "\n\t":
                var v41 string = (("hello" + "\n\t") + "")
                _ = v41
            }
        }
    } else {
        v42 := GenS1{F0: 739, F1: 1, F2: 127}
        _ = v42
    }
    v3 = v4
    var v43 uint64 = 1
    v44 := GenS3{F0: 0, F1: ("hello" <= (((("a" + (((("\n\t" + "a") + "hello") + "\n\t") + "hello")) + "\n\t") + "a") + ((("a" + "a") + "\n\t") + "\n\t")))}
    for fi45 := 0; fi45 < 17; fi45++ {
        v46 := GenS5{F0: (^(v43)), F1: v43, F2: genFn0(2147483647, ((("a" + ("hello" + (("a" + "\n\t") + ("\n\t" + "\n\t")))) + "hello") + "\n\t"))}
        _ = fi45
        _ = v46
    }
    v1 = v1
    var v47 float32 = 1.5
    var v48 uint = 18446744073709551615
    for fi49 := 0; fi49 < 16; fi49++ {
        for fi50 := 0; fi50 < 17; fi50++ {
            var v51 uint8 = 128
            v52 := []int{fi49, genFn0(998, ("\n\t" + ""))}
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi49
    }
    v43 -= (v43 / ((1) | 1))
    if true {
        v53 := GenS9{F0: 42}
        _ = v53
    } else {
        v48 += (v48 + (v48 - uint(v43)))
        v43 -= (^(v43))
        for fi54 := 0; fi54 < 9; fi54++ {
            v4 = (v4 << (((v48 ^ (v48 | (v48 - (v48 ^ v48))))) & 63))
            _ = fi54
        }
    }
    v55 := GenS10{F0: 857}
    var v56 string = "a"
    var v57 float32 = -2.25
    v2["a"] = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v48
    _ = v55
    _ = v56
    _ = v57
}
